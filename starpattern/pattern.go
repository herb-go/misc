package starpattern

import "strings"

type Options struct {
	Wildcard   rune
	Exception  []rune
	Avaliable  []rune
	NotEmpty   bool
	IgnoreCase bool
}

var Star = &Options{
	Wildcard: '*',
}

func (o *Options) parseParts(pattern string) []*part {
	var result = []*part{}
	var current = []rune{}
	for _, v := range pattern {
		if v == o.Wildcard {
			if len(current) > 0 {
				result = append(result, newPlainPart(current))
				current = []rune{}
			}
			result = append(result, wildPart)
		} else {
			current = append(current, v)
		}
	}
	if len(current) > 0 {
		result = append(result, newPlainPart(current))
	}
	return result
}
func (o *Options) New(pattern string) *Pattern {
	return &Pattern{
		options: o,
		parts:   o.parseParts(pattern),
	}
}

func (o *Options) equalRunes(a []rune, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if !o.equal(a[k], b[k]) {
			return false
		}
	}
	return true
}

func (o *Options) equal(a rune, b rune) bool {
	if o.IgnoreCase {
		return strings.ToLower(string(a)) == strings.ToLower(string(b))
	}
	return a == b
}

type part struct {
	wildcard bool
	value    []rune
}

var wildPart = &part{
	wildcard: true,
}

func newPlainPart(value []rune) *part {
	return &part{
		value: value,
	}
}

type Pattern struct {
	options *Options
	parts   []*part
}

func (p *Pattern) find(value []rune, parts []*part) (bool, []string) {
	if len(parts) == 0 && len(value) == 0 {
		return true, []string{}
	}
	if parts[0].wildcard {
		return p.findwildcard(value, parts)
	}
	return p.findplain(value, parts)
}
func (p *Pattern) findplain(value []rune, parts []*part) (bool, []string) {
	partvalue := parts[0].value
	d := len(value) - len(partvalue)
	if d == 0 {
		if p.options.equalRunes(value, parts[0].value) {
			return p.find(value[len(value):], parts[1:])
		}
	} else if d > 0 {
		if len(parts) > 1 {
			if p.options.equalRunes(value[:len(partvalue)], partvalue) {
				ok, found := p.find(value[len(partvalue):], parts[1:])
				if ok {
					return true, found
				}
			}
		}
	}
	return false, nil
}
func (p *Pattern) findwildcard(value []rune, parts []*part) (bool, []string) {
	if len(parts) == 1 {
		if p.options.NotEmpty && len(value) == 0 {
			return false, nil
		}
		return true, []string{string(value)}
	}
	if parts[1].wildcard {
		ok, found := p.findwildcard(value, parts[1:])
		if ok {
			return true, append([]string{""}, found...)
		}
	}
	if !p.options.NotEmpty {
		ok, found := p.findplain(value, parts[1:])
		if ok {
			return true, append([]string{""}, found...)
		}
	}
	var stop bool
	for k := range value {
		if len(p.options.Exception) > 0 {
			for _, e := range p.options.Exception {
				if p.options.equal(e, value[k]) {
					stop = true
				}
			}
		}
		if !stop && len(p.options.Avaliable) > 0 {
			var ok bool
			for _, a := range p.options.Avaliable {
				if p.options.equal(a, value[k]) {
					ok = true
					break
				}
			}
			if !ok {
				stop = true
			}
		}
		ok, found := p.findplain(value[k:], parts[1:])
		if ok {
			return true, append([]string{string(value[:k])}, found...)
		}
		if stop {
			break
		}
	}
	return false, nil

}
func (p *Pattern) Find(value string) (bool, []string) {
	if len(p.parts) == 0 && len(value) == 0 {
		return true, []string{}
	}
	if len(p.parts) == 1 && p.parts[0] == wildPart {
		if p.options.NotEmpty && len(value) == 0 {
			return false, nil
		}
		return true, []string{value}
	}
	data := []rune(value)
	ok, result := p.find(data, p.parts)
	if !ok {
		return false, nil
	}
	return true, append([]string{value}, result...)
}
func (p *Pattern) Match(value string) bool {
	ok, _ := p.Find(value)
	return ok
}
func New(pattern string) *Pattern {
	return Star.New(pattern)
}
