package store

//                    ObjectIndex Functions
// =======================================================================

// ObjectIndex is a table of all Objects contained by a Store.
// Object names are the map keys.  The object name is the file
// file name that stores the object data (and optionally some store
// meta data, in certain cases.
type ObjectIndex map[string]*Object

// =======================================================================
// Basic container operations
// =======================================================================

// Exists will tell if the item exists
func (s *Store) Exists(name string) bool {
	return (s.get(name) != nil)
}

// NotExists will tell if the named object does not exist.
func (s *Store) NotExists(name string) bool {
	return !s.Exists(name)
}

// Get object with name from *Store.  If object does not exist return nil
func (s *Store) get(name string) *Object {
	if o, e := s.index[name]; e {
		o.accessed++
		return o
	}
	return nil
}

// Set the object on store.  If the object does not exist it will be created,
// if the object exists, it will be overwritten.
func (s *Store) set(name string, obj *Object) {
	s.index[name] = obj
	obj.accessed++
	s.Debugf("Set object %s ++ count %d", name, len(s.index))
}
