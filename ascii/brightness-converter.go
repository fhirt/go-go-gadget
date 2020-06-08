package ascii

const characters string = "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
const maxRGB float32 = 1<<16 - 1

var StandardConverter *standardConverter = new(standardConverter)

type standardConverter struct {
}

func (p *standardConverter) ToAscii(brightness int) string {
	chars := len(characters)
	degree := float32(brightness) / maxRGB
	selectedChar := int(float32(chars) * degree)
	if selectedChar > chars-1 {
		selectedChar = chars - 1
	}

	return string(characters[selectedChar])
}
