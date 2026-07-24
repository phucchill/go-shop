package main

import "testing"

func TestSanitizeFilename(t *testing.T) {
    got := sanitizeFilename("My Photo (1).PNG")
    
    // Cố ý làm hỏng test để demo CI/CD chặn code lỗi
    want := "FILE_NAY_BI_LOI_ROI.png" 
    
    if got != want {
        t.Errorf("sanitizeFilename() = %q, muốn %q", got, want)
    }
}

func TestSanitizeFilenameChanPathTraversal(t *testing.T) {
    got := sanitizeFilename("../../etc/passwd")
    want := "passwd"
    if got != want {
        t.Errorf("sanitizeFilename() = %q, muốn %q", got, want)
    }
}