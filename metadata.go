package torrentmeta

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/anacrolix/torrent/metainfo"
)

// Metadata holds torrent metadata and provides methods to access specific fields
type Metadata struct {
	metaInfo *metainfo.MetaInfo
	info     metainfo.Info
}

// NewTorrentMetadata parses the torrent data and returns a TorrentMetadata instance
func NewTorrentMetadata(torrentData []byte) (*Metadata, error) {
	metaInfo, err := metainfo.Load(bytes.NewReader(torrentData))
	if err != nil {
		return nil, fmt.Errorf("metainfo.Load error: %v", err)
	}

	info, err := metaInfo.UnmarshalInfo()
	if err != nil {
		return nil, fmt.Errorf("metainfo.UnmarshalInfo error: %v", err)
	}

	return &Metadata{
		metaInfo: metaInfo,
		info:     info,
	}, nil
}

// Name returns the name of the torrent
func (tm *Metadata) Name() string {
	if tm.info.NameUtf8 != "" {
		return tm.info.NameUtf8
	}
	return tm.info.Name
}

// Tracker returns the tracker domain (second-level and top-level domain) extracted from the announce URL
func (tm *Metadata) Tracker() (string, error) {
	announceURL := tm.metaInfo.Announce

	// Parse the announce URL into a URL object
	parsedURL, err := url.Parse(announceURL)
	if err != nil {
		return "", fmt.Errorf("url.Parse error: %v", err)
	}

	// Extract the hostname from the announce URL
	host := parsedURL.Hostname()

	// Extract the TLD and second-level domain (e.g., "example.com" from "foo.example.com")
	parts := strings.Split(host, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid tracker domain: %s", host)
	}

	// Return the second-level domain and TLD (e.g., "example.com")
	return strings.Join(parts[len(parts)-2:], "."), nil
}
