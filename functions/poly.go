package functions

type polyNode struct {
	degree int
	coeff  float32
	next   *polyNode
}

type poly struct {
	head *polyNode
}

func pow(x float32, p int) (result float32) {
	result = 1
	for p > 0 {
		result *= x
		p--
	}
	return
}

func (p *poly) Eval(x float32) (result float32) {
	tmp := p.head
	for tmp != nil {
		result += tmp.coeff * pow(x, tmp.degree)
		tmp = tmp.next
	}
	return
}

func (p *poly) GetCoeff(deg int) (result float32) {
	tmp := p.head
	for tmp != nil {
		if tmp.degree == deg {
			result = tmp.coeff
			break
		}
		if tmp.degree < deg {
			break
		}
		tmp = tmp.next
	}
	return
}

func (p *poly) SetCoeff(deg int, c float32) {
	if p.head == nil {
		p.head = &polyNode{
			degree: deg,
			coeff:  c,
			next:   nil}
	} else {
		tmp := p.head
		for tmp.next != nil {
			if tmp.next.degree < deg {
				break
			}
			tmp = tmp.next
		}

		if tmp.next.degree == deg {
			tmp = tmp.next
		}

		if tmp.degree == deg {
			tmp.coeff = c
		} else {
			tmp.next = &polyNode{
				degree: deg,
				coeff:  c,
				next:   tmp.next,
			}
		}
	}
}

func (p *poly) Degree() int {
	if p.head == nil {
		return 0
	}
	return p.head.degree
}

func (p *poly) Clone() Polynomial {
	result := &poly{}
	if p.head != nil {
		tmp := p.head
		curr := &polyNode{
			degree: tmp.degree,
			coeff:  tmp.coeff,
			next:   nil,
		}
		result.head = curr

		tmp = tmp.next
		for tmp != nil {
			curr.next = &polyNode{
				degree: tmp.degree,
				coeff:  tmp.coeff,
				next:   nil,
			}

			curr = curr.next
			tmp = tmp.next

		}
	}
	return result
}

// NewPoly returns a new polynomial
func NewPoly() Polynomial {
	return &poly{}
}
