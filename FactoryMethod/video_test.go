package FactoryMethod

import "testing"

func TestGetVideo(t *testing.T) {

	videoFactory := NewJavaFactory()
	videoFactory.getVideo().produce()

	videoFactory2 := NewPythonFactory()
	videoFactory2.getVideo().produce()

}
