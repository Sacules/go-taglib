// Go wrapper for taglib

// Generate stringer method for types

package taglib

// #cgo pkg-config: taglib
// #cgo LDFLAGS: -ltag_c
// #include <stdlib.h>
// #include <tag_c.h>
import "C"

import (
	"errors"
	"sync"
	"time"
	"unsafe"
)

func init() {
	// Make everything utf-8
	C.taglib_id3v2_set_default_text_encoding(3)
}

type File struct {
	sync.Mutex
	fp    *C.TagLib_File
	tag   *C.TagLib_Tag
	props *C.TagLib_AudioProperties
}

// Read and parses a music file. Returns an error if the provided filename is
// not a valid file.
func Read(filename string) (*File, error) {
	m := &sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))

	fp := C.taglib_file_new(cs)
	if fp == nil || C.taglib_file_is_valid(fp) == 0 {
		return nil, errors.New("invalid file")
	}

	return &File{
		fp:    fp,
		tag:   C.taglib_file_tag(fp),
		props: C.taglib_file_audioproperties(fp),
	}, nil
}

// Close and free the file.
func (file *File) Close() {
	file.Lock()
	defer file.Unlock()

	C.taglib_file_free(file.fp)
	file.fp = nil
	file.tag = nil
	file.props = nil
}

func convertAndFree(cs *C.char) string {
	if cs == nil {
		return ""
	}

	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

// Returns a string with this tag's title.
func (file *File) Title() string {
	file.Lock()
	defer file.Unlock()

	return convertAndFree(C.taglib_tag_title(file.tag))
}

// Returns a string with this tag's artist.
func (file *File) Artist() string {
	file.Lock()
	defer file.Unlock()

	return convertAndFree(C.taglib_tag_artist(file.tag))
}

// Returns a string with this tag's album name.
func (file *File) Album() string {
	file.Lock()
	defer file.Unlock()

	return convertAndFree(C.taglib_tag_album(file.tag))
}

// Returns a string with this tag's comment.
func (file *File) Comment() string {
	file.Lock()
	defer file.Unlock()

	return convertAndFree(C.taglib_tag_comment(file.tag))
}

// Returns a string with this tag's genre.
func (file *File) Genre() string {
	file.Lock()
	defer file.Unlock()

	return convertAndFree(C.taglib_tag_genre(file.tag))
}

// Returns the tag's year or 0 if year is not set.
func (file *File) Year() int {
	file.Lock()
	defer file.Unlock()

	return int(C.taglib_tag_year(file.tag))
}

// Returns the tag's track number or 0 if track number is not set.
func (file *File) Track() int {
	file.Lock()
	defer file.Unlock()

	return int(C.taglib_tag_track(file.tag))
}

// Returns the length of the file.
func (file *File) Length() time.Duration {
	file.Lock()
	defer file.Unlock()

	length := C.taglib_audioproperties_length(file.props)
	return time.Duration(length) * time.Second
}

// Returns the bitrate of the file in kb/s.
func (file *File) Bitrate() int {
	file.Lock()
	defer file.Unlock()

	return int(C.taglib_audioproperties_bitrate(file.props))
}

// Returns the sample rate of the file in Hz.
func (file *File) Samplerate() int {
	file.Lock()
	defer file.Unlock()

	return int(C.taglib_audioproperties_samplerate(file.props))
}

// Returns the number of channels in the audio stream.
func (file *File) Channels() int {
	file.Lock()
	defer file.Unlock()

	return int(C.taglib_audioproperties_channels(file.props))
}

func init() {
	m := &sync.Mutex{}
	m.Lock()
	defer m.Unlock()

	C.taglib_set_string_management_enabled(0)
}

// Saves the \a file to disk.
func (file *File) Save() error {
	var err error
	file.Lock()
	defer file.Unlock()
	if C.taglib_file_save(file.fp) != 1 {
		err = errors.New("Cannot save file")
	}
	return err
}

// Sets the tag's title.
func (file *File) SetTitle(s string) {
	file.Lock()
	defer file.Unlock()
	cs := GetCCharPointer(s)
	defer C.free(unsafe.Pointer(cs))
	C.taglib_tag_set_title(file.tag, cs)

}

// Sets the tag's artist.
func (file *File) SetArtist(s string) {
	file.Lock()
	defer file.Unlock()
	cs := GetCCharPointer(s)
	defer C.free(unsafe.Pointer(cs))
	C.taglib_tag_set_artist(file.tag, cs)
}

// Sets the tag's album.
func (file *File) SetAlbum(s string) {
	file.Lock()
	defer file.Unlock()
	cs := GetCCharPointer(s)
	defer C.free(unsafe.Pointer(cs))
	C.taglib_tag_set_album(file.tag, cs)
}

// Sets the tag's comment.
func (file *File) SetComment(s string) {
	file.Lock()
	defer file.Unlock()
	cs := GetCCharPointer(s)
	defer C.free(unsafe.Pointer(cs))
	C.taglib_tag_set_comment(file.tag, cs)
}

// Sets the tag's genre.
func (file *File) SetGenre(s string) {
	file.Lock()
	defer file.Unlock()
	cs := GetCCharPointer(s)
	defer C.free(unsafe.Pointer(cs))
	C.taglib_tag_set_genre(file.tag, cs)
}

// Sets the tag's year.  0 indicates that this field should be cleared.
func (file *File) SetYear(i int) {
	file.Lock()
	defer file.Unlock()
	ci := C.uint(i)
	C.taglib_tag_set_year(file.tag, ci)
}

// Sets the tag's track number.  0 indicates that this field should be cleared.
func (file *File) SetTrack(i int) {
	file.Lock()
	defer file.Unlock()
	ci := C.uint(i)
	C.taglib_tag_set_track(file.tag, ci)
}

func GetCCharPointer(s string) *C.char {
	// Add a 0x00 to end
	b := append([]byte(s), 0)
	return (*C.char)(C.CBytes(b))
}
