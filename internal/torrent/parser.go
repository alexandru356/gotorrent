package torrent

import (
	"bytes"
	"errors"
	"github.com/jackpal/bencode-go"
	"os"
)

type TorrentMeta struct {
	Announce    string
	InfoHash    string
	PieceHashes []string
	PieceLength int64
	Length      int
	Name        string
}

func ParseTorrent(filePath string) (*TorrentMeta, error) {

	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var torrent map[string]interface{}
	err = bencode.Unmarshal(bytes.NewReader(raw), &torrent)
	if err != nil {
		return nil, err
	}

	info, ok := torrent["info"].(map[string]interface{})
	if !ok {
		return nil, errors.New("info dict missing")
	}

	var infoBuffer bytes.Buffer
	err = bencode.Marshal(&infoBuffer, info)
	if err != nil {
		return nil, err
	}

	infoHash := sha1.Sum(infoBuffer.Bytes())

}
