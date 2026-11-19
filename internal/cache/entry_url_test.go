package cache

import "testing"

func TestPyPyYouTubeFallback(t *testing.T) {
	entry, ok := NewEntry("pypy_3612").(*UrlBasedEntry)
	if !ok {
		t.Fatal("expected pypy_3612 to be a UrlBasedEntry")
	}
	if entry.youtubeFallback == nil {
		t.Fatal("expected a PyPyDance entry to have a YouTube fallback")
	}

	got, ok := entry.youtubeFallback("iBJmDt1lxE0")
	want := "https://cdn.pypy.dance/iBJmDt1lxE0.mp4"
	if !ok || got != want {
		t.Fatalf("unexpected fallback url: %q (ok=%v), want %q", got, ok, want)
	}
}

func TestYouTubeFallbackOnlyForPyPy(t *testing.T) {
	for _, id := range []string{"wanna_1", "dudu_1", "bili_BV17g7XzME13", "yt_dQw4w9WgXcQ"} {
		entry, ok := NewEntry(id).(*UrlBasedEntry)
		if !ok {
			t.Fatalf("expected %s to be a UrlBasedEntry", id)
		}
		if entry.youtubeFallback != nil {
			t.Fatalf("expected %s to have no YouTube fallback", id)
		}
	}
}
