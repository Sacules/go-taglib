package taglib

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

func TestReadNothing(t *testing.T) {
	file, err := Read("doesnotexist.mp3")

	if file != nil {
		t.Fatal("Returned non nil file struct.")
	}

	if err == nil {
		t.Fatal("Returned nil err")
	}
}

func TestReadDirectory(t *testing.T) {
	file, err := Read("/")

	if file != nil {
		t.Fatal("Returned non nil file struct.")
	}

	if err == nil {
		t.Fatal("Returned nil err.")
	}
}

func TestTagLib(t *testing.T) {
	// MP3
	filemp3, err := Read("test.mp3")
	defer filemp3.Close()

	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := filemp3.Title(); title != "The Title" {
		t.Errorf("Got wrong title: %s", title)
	}

	if artist := filemp3.Artist(); artist != "The Artist" {
		t.Errorf("Got wrong artist: %s", artist)
	}

	if album := filemp3.Album(); album != "The Album" {
		t.Errorf("Got wrong album: %s", album)
	}

	if comment := filemp3.Comment(); comment != "A Comment" {
		t.Errorf("Got wrong comment: %s", comment)
	}

	if genre := filemp3.Genre(); genre != "Booty Bass" {
		t.Errorf("Got wrong genre: %s", genre)
	}

	if year := filemp3.Year(); year != 1942 {
		t.Errorf("Got wrong year: %d", year)
	}

	if track := filemp3.Track(); track != 42 {
		t.Errorf("Got wrong track: %d", track)
	}

	// Test the properties
	if length := filemp3.Length(); length != 42*time.Second {
		t.Errorf("Got wrong length: %s", length)
	}

	if bitrate := filemp3.Bitrate(); bitrate != 128 {
		t.Errorf("Got wrong bitrate: %d", bitrate)
	}

	if samplerate := filemp3.Samplerate(); samplerate != 44100 {
		t.Errorf("Got wrong samplerate: %d", samplerate)
	}

	if channels := filemp3.Channels(); channels != 2 {
		t.Errorf("Got wrong channels: %d", channels)
	}
}

func TestTagLibOGG(t *testing.T) {
	// OGG
	fileogg, err := Read("test.ogg")
	defer fileogg.Close()

	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := fileogg.Title(); title != "Free Software Song" {
		t.Errorf("Got wrong title: %s", title)
	}

	if artist := fileogg.Artist(); artist != "Mark Forry, Yvette Osborne, Ron Fox, Steve Finney, Bill Cope, Kip McAtee, Ernie Provencher, Dan Auvil" {
		t.Errorf("Got wrong artist: %s", artist)
	}

	if album := fileogg.Album(); album != "Freedom" {
		t.Errorf("Got wrong album: %s", album)
	}

	if comment := fileogg.Comment(); comment != "" {
		t.Errorf("Got wrong comment: %s", comment)
	}

	if genre := fileogg.Genre(); genre != "Ethnic" {
		t.Errorf("Got wrong genre: %s", genre)
	}

	if year := fileogg.Year(); year != 2009 {
		t.Errorf("Got wrong year: %d", year)
	}

	if track := fileogg.Track(); track != 1 {
		t.Errorf("Got wrong track: %d", track)
	}

	// Test the properties
	if length := fileogg.Length(); length != 10*time.Second {
		t.Errorf("Got wrong length: %s", length)
	}

	if bitrate := fileogg.Bitrate(); bitrate != 153 {
		t.Errorf("Got wrong bitrate: %d", bitrate)
	}

	if samplerate := fileogg.Samplerate(); samplerate != 44100 {
		t.Errorf("Got wrong samplerate: %d", samplerate)
	}

	if channels := fileogg.Channels(); channels != 2 {
		t.Errorf("Got wrong channels: %d", channels)
	}
}

func TestTagLibFLAC(t *testing.T) {
	// FLAC
	fileflac, err := Read("test.flac")
	defer fileflac.Close()

	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := fileflac.Title(); title != "Free Software Song" {
		t.Errorf("Got wrong title: %s", title)
	}

	if artist := fileflac.Artist(); artist != "Mark Forry, Yvette Osborne, Ron Fox, Steve Finney, Bill Cope, Kip McAtee, Ernie Provencher, Dan Auvil" {
		t.Errorf("Got wrong artist: %s", artist)
	}

	if album := fileflac.Album(); album != "Freedom" {
		t.Errorf("Got wrong album: %s", album)
	}

	if comment := fileflac.Comment(); comment != "" {
		t.Errorf("Got wrong comment: %s", comment)
	}

	if genre := fileflac.Genre(); genre != "Ethnic" {
		t.Errorf("Got wrong genre: %s", genre)
	}

	if year := fileflac.Year(); year != 2009 {
		t.Errorf("Got wrong year: %d", year)
	}

	if track := fileflac.Track(); track != 1 {
		t.Errorf("Got wrong track: %d", track)
	}

	// Test the properties
	if length := fileflac.Length(); length != 10*time.Second {
		t.Errorf("Got wrong length: %s", length)
	}

	if bitrate := fileflac.Bitrate(); bitrate != 816 {
		t.Errorf("Got wrong bitrate: %d", bitrate)
	}

	if samplerate := fileflac.Samplerate(); samplerate != 44100 {
		t.Errorf("Got wrong samplerate: %d", samplerate)
	}

	if channels := fileflac.Channels(); channels != 2 {
		t.Errorf("Got wrong channels: %d", channels)
	}
}

func TestWriteTagLibMP3(t *testing.T) {
	fileName := "test.mp3"
	file, err := Read(fileName)
	defer file.Close()
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	tempDir, err := ioutil.TempDir("", "go-taglib-test-MP3")
	if err != nil {
		t.Fatalf("Cannot create temporary file for writing tests: %s", err)
	}

	tempFileName := path.Join(tempDir, "go-taglib-test.mp3")
	defer os.RemoveAll(tempDir)

	err = cp(tempFileName, fileName)
	if err != nil {
		t.Fatalf("Cannot copy file for writing tests: %s", err)
	}

	modifiedFile, err := Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	modifiedFile.SetAlbum(getModifiedString(file.Album()))
	modifiedFile.SetComment(getModifiedString(file.Comment()))
	modifiedFile.SetGenre(getModifiedString(file.Genre()))
	modifiedFile.SetTrack(file.Track() + 1)
	modifiedFile.SetYear(file.Year() + 1)
	modifiedFile.SetArtist(getModifiedString(file.Artist()))
	modifiedFile.SetTitle(getModifiedString(file.Title()))

	err = modifiedFile.Save()
	if err != nil {
		t.Fatalf("Cannot save file : %s", err)
	}

	modifiedFile.Close()

	//Re-open the modified file
	modifiedFile, err = Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := modifiedFile.Title(); title != getModifiedString("The Title") {
		t.Errorf("Got wrong modified title: %s", title)
	}

	if artist := modifiedFile.Artist(); artist != getModifiedString("The Artist") {
		t.Errorf("Got wrong modified artist: %s", artist)
	}

	if album := modifiedFile.Album(); album != getModifiedString("The Album") {
		t.Errorf("Got wrong modified album: %s", album)
	}

	if comment := modifiedFile.Comment(); comment != getModifiedString("A Comment") {
		t.Errorf("Got wrong modified comment: %s", comment)
	}

	if genre := modifiedFile.Genre(); genre != getModifiedString("Booty Bass") {
		t.Errorf("Got wrong modified genre: %s", genre)
	}

	if year := modifiedFile.Year(); year != getModifiedInt(1942) {
		t.Errorf("Got wrong modified year: %d", year)
	}

	if track := modifiedFile.Track(); track != getModifiedInt(42) {
		t.Errorf("Got wrong modified track: %d", track)
	}
}

func TestWriteTagLibOGG(t *testing.T) {
	fileName := "test.ogg"
	file, err := Read(fileName)
	defer file.Close()
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	tempDir, err := ioutil.TempDir("", "go-taglib-test-OGG")
	if err != nil {
		t.Fatalf("Cannot create temporary file for writing tests: %s", err)
	}

	tempFileName := path.Join(tempDir, "go-taglib-test.ogg")
	defer os.RemoveAll(tempDir)

	err = cp(tempFileName, fileName)
	if err != nil {
		t.Fatalf("Cannot copy file for writing tests: %s", err)
	}

	modifiedFile, err := Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	modifiedFile.SetAlbum(getModifiedString(file.Album()))
	modifiedFile.SetComment(getModifiedString(file.Comment()))
	modifiedFile.SetGenre(getModifiedString(file.Genre()))
	modifiedFile.SetTrack(file.Track() + 1)
	modifiedFile.SetYear(file.Year() + 1)
	modifiedFile.SetArtist(getModifiedString(file.Artist()))
	modifiedFile.SetTitle(getModifiedString(file.Title()))

	err = modifiedFile.Save()
	if err != nil {
		t.Fatalf("Cannot save file : %s", err)
	}

	modifiedFile.Close()

	//Re-open the modified file
	modifiedFile, err = Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := modifiedFile.Title(); title != getModifiedString("Free Software Song") {
		t.Errorf("Got wrong modified title: %s", title)
	}

	if artist := modifiedFile.Artist(); artist != getModifiedString("Mark Forry, Yvette Osborne, Ron Fox, Steve Finney, Bill Cope, Kip McAtee, Ernie Provencher, Dan Auvil") {
		t.Errorf("Got wrong modified artist: %s", artist)
	}

	if album := modifiedFile.Album(); album != getModifiedString("Freedom") {
		t.Errorf("Got wrong modified album: %s", album)
	}

	if comment := modifiedFile.Comment(); comment != getModifiedString("") {
		t.Errorf("Got wrong modified comment: %s", comment)
	}

	if genre := modifiedFile.Genre(); genre != getModifiedString("Ethnic") {
		t.Errorf("Got wrong modified genre: %s", genre)
	}

	if year := modifiedFile.Year(); year != getModifiedInt(2009) {
		t.Errorf("Got wrong modified year: %d", year)
	}

	if track := modifiedFile.Track(); track != getModifiedInt(1) {
		t.Errorf("Got wrong modified track: %d", track)
	}
}

func TestGenericWriteTagLib(t *testing.T) {
	fileName := "test.mp3"
	file, err := Read(fileName)
	defer file.Close()

	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}
	tempDir, err := ioutil.TempDir("", "go-taglib-test")

	if err != nil {
		t.Fatalf("Cannot create temporary file for writing tests: %s", err)
	}

	tempFileName := path.Join(tempDir, "go-taglib-test.mp3")

	defer os.RemoveAll(tempDir)

	err = cp(tempFileName, fileName)

	if err != nil {
		t.Fatalf("Cannot copy file for writing tests: %s", err)
	}

	modifiedFile, err := Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}
	modifiedFile.SetAlbum(getModifiedString(file.Album()))
	modifiedFile.SetComment(getModifiedString(file.Comment()))
	modifiedFile.SetGenre(getModifiedString(file.Genre()))
	modifiedFile.SetTrack(file.Track() + 1)
	modifiedFile.SetYear(file.Year() + 1)
	modifiedFile.SetArtist(getModifiedString(file.Artist()))
	modifiedFile.SetTitle(getModifiedString(file.Title()))
	err = modifiedFile.Save()
	if err != nil {
		t.Fatalf("Cannot save file : %s", err)
	}
	modifiedFile.Close()
	//Re-open the modified file
	modifiedFile, err = Read(tempFileName)
	if err != nil {
		t.Fatalf("Read returned error: %s", err)
	}

	// Test the Tags
	if title := modifiedFile.Title(); title != getModifiedString("The Title") {
		t.Errorf("Got wrong modified title: %s", title)
	}

	if artist := modifiedFile.Artist(); artist != getModifiedString("The Artist") {
		t.Errorf("Got wrong modified artist: %s", artist)
	}

	if album := modifiedFile.Album(); album != getModifiedString("The Album") {
		t.Errorf("Got wrong modified album: %s", album)
	}

	if comment := modifiedFile.Comment(); comment != getModifiedString("A Comment") {
		t.Errorf("Got wrong modified comment: %s", comment)
	}

	if genre := modifiedFile.Genre(); genre != getModifiedString("Booty Bass") {
		t.Errorf("Got wrong modified genre: %s", genre)
	}

	if year := modifiedFile.Year(); year != getModifiedInt(1942) {
		t.Errorf("Got wrong modified year: %d", year)
	}

	if track := modifiedFile.Track(); track != getModifiedInt(42) {
		t.Errorf("Got wrong modified track: %d", track)
	}
}

func checkModified(original string, modified string) bool {
	return modified == getModifiedString(original)
}

func getModifiedString(s string) string {
	return s + " MODIFIED"
}

func getModifiedInt(i int) int {
	return i + 1
}

func cp(dst, src string) error {
	s, err := os.Open(src)
	defer s.Close()
	if err != nil {
		return err
	}
	// no need to check errors on read only file, we already got everything
	// we need from the filesystem, so nothing can go wrong now.
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
