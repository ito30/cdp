package pizza

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPizzaToppings(t *testing.T) {
	tests := map[string]struct {
		pizza func() Pizza
	}{
		"veggie pizza": {
			pizza: func() Pizza {
				return &VeggieMania{}
			},
		},
		"veggie with cheese pizza": {
			pizza: func() Pizza {
				veggie := &VeggieMania{}
				return &CheeseTopping{pizza: veggie}
			},
		},
		"veggie with cheese & tomato pizza": {
			pizza: func() Pizza {
				veggie := &VeggieMania{}
				veggieWithCheese := &CheeseTopping{pizza: veggie}
				return &TomatoTopping{pizza: veggieWithCheese}
			},
		},
		"cheese topping only": {
			pizza: func() Pizza {
				return &CheeseTopping{}
			},
		},
	}

	for name, tt := range tests {
		fmt.Printf("Price of %s is %d\n", name, tt.pizza().getPrice())
	}
}

func TestPizza(t *testing.T) {
	pizza := &VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	total := pizzaWithCheeseAndTomato.getPrice()
	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", total)

	assert.Equal(t, 32, total)
}
