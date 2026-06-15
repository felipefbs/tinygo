package display

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

type Display struct {
	display *ssd1306.Device
}

func (d Display) PrintMessage(message string) {
	// Limpa o buffer anterior para as letras não se sobreporem
	d.display.ClearDisplay()

	textColor := color.RGBA{255, 255, 255, 255}

	// Escreve a nova mensagem nas coordenadas
	tinyfont.WriteLine(d.display, &proggy.TinySZ8pt7b, 28, 40, message, textColor)

	// Envia o buffer atualizado para a tela física
	d.display.Display()
}

func InitDisplay() *Display {
	machine.I2C0.Configure(machine.I2CConfig{
		SCL:       machine.GPIO6,
		SDA:       machine.GPIO5,
		Frequency: 400 * machine.KHz,
	})

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()
	display.ClearBuffer()

	return &Display{
		display: display,
	}
}
