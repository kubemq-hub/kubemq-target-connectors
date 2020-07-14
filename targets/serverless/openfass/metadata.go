package openfass

import (
	"fmt"
	"github.com/kubemq-hub/kubemq-target-connectors/types"
)

type metadata struct {
	topic string
}

func parseMetadata(meta types.Metadata) (metadata, error) {
	m := metadata{}
	var err error
	m.topic, err = meta.MustParseString("topic")
	if err != nil {
		return metadata{}, fmt.Errorf("error parsing topic func, %w", err)
	}
	return m, nil
}
