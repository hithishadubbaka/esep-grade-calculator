package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 35, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 20, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}



func TestGetGradeD(t *testing.T) {
	expected := "D"
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 65, Assignment)
	gc.AddGrade("exam", 60, Exam)
	gc.AddGrade("essay", 61, Essay)

	if gc.GetFinalGrade() != expected {
		t.Errorf("Expected %s, got %s", expected, gc.GetFinalGrade())
	}
}



func TestBoundaryCtoD(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 70, Assignment)
	gc.AddGrade("exam", 70, Exam)
	gc.AddGrade("essay", 70, Essay)

	if gc.GetFinalGrade() != "C" {
		t.Errorf("Expected C, got %s", gc.GetFinalGrade())
	}
}



func TestBoundaryBtoA(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment", 89, Assignment)
	gc.AddGrade("exam", 89, Exam)
	gc.AddGrade("essay", 89, Essay)

	if gc.GetFinalGrade() != "B" {
		t.Errorf("Expected B, got %s", gc.GetFinalGrade())
	}
}



func TestEmptyGrades(t *testing.T) {
	gc := NewGradeCalculator()
	if gc.GetFinalGrade() != "F" {
		t.Errorf("Expected F for no grades, got %s", gc.GetFinalGrade())
	}
}



func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment', got %s", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam', got %s", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay', got %s", Essay.String())
	}
}

