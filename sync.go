package textsecure

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/prometheus/common/log"
	"github.com/signal-golang/textsecure/protobuf"
)

// handleSyncMessage handles an incoming SyncMessage.
func handleSyncMessage(src string, timestamp uint64, sm *signalservice.SyncMessage) error {
	log.Debugf("[textsecure] SyncMessage recieved at %d", timestamp)

	if sm.GetSent() != nil {
		log.Debugln("[textsecure] SyncMessage getSent")
		return handleSyncSent(sm.GetSent(), timestamp)
	} else if sm.GetContacts() != nil {
		log.Debugln("[textsecure] SyncMessage contacts")
		return nil
	} else if sm.GetGroups() != nil {
		log.Debugln("[textsecure] SyncMessage groups")
		return nil
	} else if sm.GetRequest() != nil {
		log.Debugln("[textsecure] SyncMessage getRequest")
		return handleSyncRequest(sm.GetRequest())
	} else if sm.GetRead() != nil {
		log.Debugln("[textsecure] SyncMessage getRead")
		return handleSyncRead(sm.GetRead())
	} else if sm.GetBlocked() != nil {
		log.Debugln("[textsecure] SyncMessage blocked")
		return nil
	} else if sm.GetVerified() != nil {
		log.Debugln("[textsecure] SyncMessage verified")
		return nil
	} else if sm.GetConfiguration() != nil {
		log.Debugln("[textsecure] SyncMessage configuration")
		return nil
	} else if sm.GetPadding() != nil {
		log.Debugln("[textsecure] SyncMessage padding")
		return nil
	} else if sm.GetStickerPackOperation() != nil {
		log.Debugln("[textsecure] SyncMessage GetStickerPackOperation")
		return nil
	} else if sm.GetViewOnceOpen() != nil {
		log.Debugln("[textsecure] SyncMessage GetViewOnceOpen")
		return nil
	} else if sm.GetFetchLatest() != nil {
		log.Debugln("[textsecure] SyncMessage GetFetchLatest")
		return nil
	} else {
		log.Errorf("[textsecure] SyncMessage contains no known sync types")
	}

	return nil
}

// handleSyncSent handles sync sent messages
func handleSyncSent(s *signalservice.SyncMessage_Sent, ts uint64) error {
	dm := s.GetMessage()
	dest := s.GetDestinationE164()
	timestamp := s.GetTimestamp()

	if dm == nil {
		return fmt.Errorf("DataMessage was nil for SyncMessage_Sent")
	}

	flags, err := handleFlags(dest, dm)
	if err != nil {
		return err
	}

	atts, err := handleAttachments(dm)
	if err != nil {
		return err
	}

	gr, err := handleGroups(dest, dm)
	if err != nil {
		return err
	}
	cs, err := handleContacts(dest, dm)
	if err != nil {
		return err
	}

	msg := &Message{
		source:      dest,
		message:     dm.GetBody(),
		attachments: atts,
		group:       gr,
		contact:     cs,
		timestamp:   timestamp,
		flags:       flags,
	}

	if client.SyncSentHandler != nil {
		client.SyncSentHandler(msg, ts)
	}

	return nil
}

// handleSyncRequestMessage
func handleSyncRequest(request *signalservice.SyncMessage_Request) error {
	if request.GetType() == signalservice.SyncMessage_Request_CONTACTS {
		return sendContactUpdate()
	} else if request.GetType() == signalservice.SyncMessage_Request_GROUPS {
		return sendGroupUpdate()
	} else {
		log.Debugln("[textsecure] handle sync request unhandled type", request.GetType())
	}

	return nil
}

// sendContactUpdate
func sendContactUpdate() error {
	log.Debugf("[textsecure] Sending contact SyncMessage")

	lc, err := GetRegisteredContacts()
	if err != nil {
		return fmt.Errorf("could not get local contacts: %s", err)
	}

	var buf bytes.Buffer

	for _, c := range lc {
		cd := &signalservice.ContactDetails{
			Number:      &c.Tel,
			Uuid:        &c.Uuid,
			Name:        &c.Name,
			Color:       &c.Color,
			Blocked:     &c.Blocked,
			ExpireTimer: &c.ExpireTimer,

			// TODO: handle avatars
		}

		b, err := proto.Marshal(cd)
		if err != nil {
			log.Errorf("[textsecure] Failed to marshal contact details")
			continue
		}

		buf.Write(varint32(len(b)))
		buf.Write(b)
	}

	a, err := uploadAttachment(&buf, "application/octet-stream")
	if err != nil {
		return err
	}

	sm := &signalservice.SyncMessage{
		Contacts: &signalservice.SyncMessage_Contacts{
			Blob: &signalservice.AttachmentPointer{
				AttachmentIdentifier: &signalservice.AttachmentPointer_CdnId{
					CdnId: a.id,
				},
				ContentType: &a.ct,
				Key:         a.keys[:],
				Digest:      a.digest[:],
				Size:        &a.size,
			},
		},
	}
	_, err = sendSyncMessage(sm, nil)
	return err
}

// sendGroupUpdate
func sendGroupUpdate() error {
	log.Debugf("Sending group SyncMessage")

	var buf bytes.Buffer

	for _, g := range groups {
		gd := &signalservice.GroupDetails{
			Id:          g.ID,
			Name:        &g.Name,
			MembersE164: g.Members,
			// XXX add support for avatar
			// XXX add support for active?
		}

		b, err := proto.Marshal(gd)
		if err != nil {
			log.Errorf("Failed to marshal group details")
			continue
		}

		buf.Write(varint32(len(b)))
		buf.Write(b)
	}

	a, err := uploadAttachment(&buf, "application/octet-stream")
	if err != nil {
		return err
	}

	sm := &signalservice.SyncMessage{
		Groups: &signalservice.SyncMessage_Groups{
			Blob: &signalservice.AttachmentPointer{
				AttachmentIdentifier: &signalservice.AttachmentPointer_CdnId{
					CdnId: a.id,
				},
				ContentType: &a.ct,
				Key:         a.keys[:],
				Digest:      a.digest[:],
				Size:        &a.size,
			},
		},
	}
	_, err = sendSyncMessage(sm, nil)
	return err
}

func handleSyncRead(readMessages []*signalservice.SyncMessage_Read) error {
	if client.SyncReadHandler != nil {
		for _, s := range readMessages {
			client.SyncReadHandler(s.GetSenderE164(), s.GetTimestamp())
		}
	}

	return nil
}

// Encodes a 32bit base 128 variable-length integer and returns the bytes
func varint32(value int) []byte {
	buf := make([]byte, binary.MaxVarintLen32)
	n := binary.PutUvarint(buf, uint64(value))
	return buf[:n]
}
