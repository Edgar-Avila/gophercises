package story

import (
	"encoding/json"
	"os"
)

type Story struct {
    arcs map[string]StoryArc
}

func (s *Story) Load (filename string) error {
    file, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    if err = json.Unmarshal(file, &s.arcs); err != nil {
        return err
    }
    return nil
}

func (s *Story) GetArc (arcname string) (StoryArc, bool) {
    arc, ok := s.arcs[arcname]
    return arc, ok
}
