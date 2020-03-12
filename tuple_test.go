package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

var tuple *Tuple

func aTuple(x, y, z, w float64) error {
	tuple = &Tuple{x, y, z, w}
	return nil
}

func aVector(x, y, z float64) error {
	tuple = NewVector(x, y, z)
	return nil
}

func aPoint(x, y, z float64) error {
	tuple = NewPoint(x, y, z)
	return nil
}

func ax(x float64) error {
	if tuple.X != x {
		return fmt.Errorf("expected a.x to equal %f but was %f", x, tuple.X)
	}
	return nil
}

func ay(expected float64) error {
	if tuple.Y != expected {
		return fmt.Errorf("expected a.y to equal %f but was %f", expected, tuple.Y)
	}
	return nil
}

func az(expected float64) error {
	if tuple.Z != expected {
		return fmt.Errorf("expected a.z to equal %f but was %f", expected, tuple.Z)
	}
	return nil
}

func aw(expected float64) error {
	if tuple.W != expected {
		return fmt.Errorf("expected a.w to equal %f but was %f", expected, tuple.W)
	}
	return nil
}

func aIsAPoint() error {
	if !tuple.IsPoint() {
		return fmt.Errorf("expected a to be a Point but it was not: %v", tuple)
	}
	return nil
}

func aIsNotAVector() error {
	if tuple.IsVector() {
		return fmt.Errorf("expected a to not be a Vector but it was: %v", tuple)
	}
	return nil
}

func aIsNotAPoint() error {
	if tuple.IsPoint() {
		return fmt.Errorf("expected a to not be a Point but it was: %v", tuple)
	}
	return nil
}

func aIsAVector() error {
	if !tuple.IsVector() {
		return fmt.Errorf("expected a to be a Vector but it was not: %v", tuple)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a ← tuple\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aTuple)
	s.Step(`^a ← vector\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aVector)
	s.Step(`^a ← point\((\-*\d+\.\d+), (\-*\d+\.\d+), (\-*\d+\.\d+)\)$`, aPoint)
	s.Step(`^a\.x = (\-*\d+\.\d+)$`, ax)
	s.Step(`^a\.y = (\-*\d+\.\d+)$`, ay)
	s.Step(`^a\.z = (\-*\d+\.\d+)$`, az)
	s.Step(`^a\.w = (\-*\d+\.\d+)$`, aw)
	s.Step(`^a is a point$`, aIsAPoint)
	s.Step(`^a is not a vector$`, aIsNotAVector)
	s.Step(`^a is not a point$`, aIsNotAPoint)
	s.Step(`^a is a vector$`, aIsAVector)
	s.BeforeScenario(func(i interface{}) {
		tuple = nil // clean the state before every scenario
	})
}
