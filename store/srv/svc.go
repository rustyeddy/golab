package store

// Start storage
func (s *Store) Start() error { return nil }

// Stop Storage
func (s *Store) Stop() error { return nil }

// Status storage status
func (s *Store) Status() string { return "Doin alright\n" }
