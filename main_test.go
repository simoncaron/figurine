// Copyright 2016 Arsham Shirvani <arshamshirvani@gmail.com>. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license
// License that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	_ "github.com/arsham/figurine/statik"
	"github.com/arsham/rainbow/rainbow"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/rakyll/statik/fs"
)

func BenchmarkGenerationPart(b *testing.B) {
	bcs := []string{
		"Arsham",
		"hRARbnf730ObNA1",
		"ZvooVEF2UOEg7k ha3IPoD319z9rWUEOUIH",
		"KjV8HeLaSV0MDiZFyXAg2XDCC MZv9O5d 1Z86mJ qw2d7Z0CAT7MrAunZH V74YD omlrSwpjXY2SxS6",
	}
	fs, err := fs.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bc := range bcs {
		name := fmt.Sprintf("%d", len(bc))
		b.Run(name, func(b *testing.B) {
			font, err := fs.Open("/" + fontNames[200])
			if err != nil {
				b.Fatal(err)
			}
			myFigure := figure.NewFigureWithFont(bc, font, true)
			seed := int(rand.Int31n(256))
			for i := 0; i < b.N; i++ {
				buf := new(bytes.Buffer)
				figure.Write(buf, myFigure)

				rb := rainbow.Light{
					Reader: buf,
					Writer: ioutil.Discard,
					Seed:   seed,
				}
				rb.Paint()
			}
		})
	}
}

func BenchmarkWholeAction(b *testing.B) {
	bcs := []string{
		"Arsham",
		"hRARbnf730ObNA1",
		"ZvooVEF2UOEg7k ha3IPoD319z9rWUEOUIH",
		"KjV8HeLaSV0MDiZFyXAg2XDCC MZv9O5d 1Z86mJ qw2d7Z0CAT7MrAunZH V74YD omlrSwpjXY2SxS6",
	}
	for _, bc := range bcs {
		name := fmt.Sprintf("%d", len(bc))
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fs, err := fs.New()
				if err != nil {
					fmt.Println(err)
					return
				}
				font, err := fs.Open("/" + fontNames[200])
				if err != nil {
					b.Fatal(err)
				}
				myFigure := figure.NewFigureWithFont(bc, font, true)
				buf := new(bytes.Buffer)
				figure.Write(buf, myFigure)

				seed := int(rand.Int31n(256))
				rb := rainbow.Light{
					Reader: buf,
					Writer: ioutil.Discard,
					Seed:   seed,
				}
				rb.Paint()
			}
		})
	}
}

var fontNames = []string{
	"1Row.flf",
	"3-D.flf",
	"3D Diagonal.flf",
	"3D-ASCII.flf",
	"3d.flf",
	"3d_diagonal.flf",
	"3x5.flf",
	"4Max.flf",
	"5 Line Oblique.flf",
	"5lineoblique.flf",
	"AMC 3 Line.flf",
	"AMC 3 Liv1.flf",
	"AMC AAA01.flf",
	"AMC Neko.flf",
	"AMC Razor.flf",
	"AMC Razor2.flf",
	"AMC Slash.flf",
	"AMC Slider.flf",
	"AMC Thin.flf",
	"AMC Tubes.flf",
	"AMC Untitled.flf",
	"ANSI Regular.flf",
	"ASCII New Roman.flf",
	"Acrobatic.flf",
	"Alligator.flf",
	"Alligator2.flf",
	"Alpha.flf",
	"Alphabet.flf",
	"Arrows.flf",
	"Avatar.flf",
	"B1FF.flf",
	"Banner.flf",
	"Banner3-D.flf",
	"Banner3.flf",
	"Banner4.flf",
	"Barbwire.flf",
	"Basic.flf",
	"Bear.flf",
	"Bell.flf",
	"Benjamin.flf",
	"Big Chief.flf",
	"Big.flf",
	"Bigfig.flf",
	"Binary.flf",
	"Block.flf",
	"Blocks.flf",
	"Bloody.flf",
	"Bolger.flf",
	"Braced.flf",
	"Bright.flf",
	"Broadway KB.flf",
	"Broadway.flf",
	"Bubble.flf",
	"Bulbhead.flf",
	"Caligraphy.flf",
	"Caligraphy2.flf",
	"Calvin S.flf",
	"Catwalk.flf",
	"Chiseled.flf",
	"Chunky.flf",
	"Coinstak.flf",
	"Cola.flf",
	"Colossal.flf",
	"Computer.flf",
	"Contessa.flf",
	"Contrast.flf",
	"Cosmike.flf",
	"Crawford.flf",
	"Crawford2.flf",
	"Crazy.flf",
	"Cricket.flf",
	"Cursive.flf",
	"Cyberlarge.flf",
	"Cybermedium.flf",
	"Cybersmall.flf",
	"Cygnet.flf",
	"DANC4.flf",
	"DOS Rebel.flf",
	"Dancing Font.flf",
	"Decimal.flf",
	"Def Leppard.flf",
	"Delta Corps Priest 1.flf",
	"Diamond.flf",
	"Diet Cola.flf",
	"Digital.flf",
	"Doh.flf",
	"Doom.flf",
	"Dot Matrix.flf",
	"Double Shorts.flf",
	"Double.flf",
	"Dr Pepper.flf",
	"Efti Chess.flf",
	"Efti Font.flf",
	"Efti Italic.flf",
	"Efti Piti.flf",
	"Efti Robot.flf",
	"Efti Wall.flf",
	"Efti Water.flf",
	"Electronic.flf",
	"Elite.flf",
	"Epic.flf",
	"Fender.flf",
	"Filter.flf",
	"Fire Font-k.flf",
	"Fire Font-s.flf",
	"Flipped.flf",
	"Flower Power.flf",
	"Four Tops.flf",
	"Fun Face.flf",
	"Fun Faces.flf",
	"Fuzzy.flf",
	"Georgi16.flf",
	"Georgia11.flf",
	"Ghost.flf",
	"Ghoulish.flf",
	"Glenyn.flf",
	"Goofy.flf",
	"Gothic.flf",
	"Graceful.flf",
	"Gradient.flf",
	"Graffiti.flf",
	"Greek.flf",
	"Henry 3D.flf",
	"Hieroglyphs.flf",
	"Hollywood.flf",
	"Horizontal Left.flf",
	"Horizontal Right.flf",
	"Impossible.flf",
	"Invita.flf",
	"Isometric1.flf",
	"Isometric2.flf",
	"Isometric3.flf",
	"Isometric4.flf",
	"Italic.flf",
	"Ivrit.flf",
	"JS Block Letters.flf",
	"JS Bracket Letters.flf",
	"JS Capital Curves.flf",
	"JS Cursive.flf",
	"JS Stick Letters.flf",
	"Jacky.flf",
	"Jazmine.flf",
	"Jerusalem.flf",
	"Kban.flf",
	"Keyboard.flf",
	"Knob.flf",
	"Konto Slant.flf",
	"Konto.flf",
	"LCD.flf",
	"Larry 3D 2.flf",
	"Larry 3D.flf",
	"Lean.flf",
	"Letters.flf",
	"Lil Devil.flf",
	"Line Blocks.flf",
	"Linux.flf",
	"Lockergnome.flf",
	"Madrid.flf",
	"Marquee.flf",
	"Maxfour.flf",
	"Merlin1.flf",
	"Merlin2.flf",
	"Mike.flf",
	"Mini.flf",
	"Mirror.flf",
	"Mnemonic.flf",
	"Modular.flf",
	"Morse.flf",
	"Moscow.flf",
	"Mshebrew210.flf",
	"Muzzle.flf",
	"NScript.flf",
	"NT Greek.flf",
	"NV Script.flf",
	"Nancyj-Fancy.flf",
	"Nancyj-Improved.flf",
	"Nancyj-Underlined.flf",
	"Nancyj.flf",
	"Nipples.flf",
	"O8.flf",
	"OS2.flf",
	"Octal.flf",
	"Ogre.flf",
	"Old Banner.flf",
	"Patorjk's Cheese.flf",
	"Patorjk-HeX.flf",
	"Pawp.flf",
	"Peaks Slant.flf",
	"Peaks.flf",
	"Pebbles.flf",
	"Pepper.flf",
	"Poison.flf",
	"Puffy.flf",
	"Puzzle.flf",
	"Pyramid.flf",
	"Rammstein.flf",
	"Rectangles.flf",
	"Red Phoenix.flf",
	"Relief.flf",
	"Relief2.flf",
	"Reverse.flf",
	"Roman.flf",
	"Rot13.flf",
	"Rotated.flf",
	"Rounded.flf",
	"Rowan Cap.flf",
	"Rozzo.flf",
	"Runic.flf",
	"Runyc.flf",
	"S Blood.flf",
	"SL Script.flf",
	"Santa Clara.flf",
	"Script.flf",
	"Serifcap.flf",
	"Shadow.flf",
	"Shimrod.flf",
	"Short.flf",
	"Slant Relief.flf",
	"Slant.flf",
	"Slide.flf",
	"Small Caps.flf",
	"Small Isometric1.flf",
	"Small Keyboard.flf",
	"Small Poison.flf",
	"Small Script.flf",
	"Small Shadow.flf",
	"Small Slant.flf",
	"Small Tengwar.flf",
	"Small.flf",
	"Soft.flf",
	"Speed.flf",
	"Spliff.flf",
	"Stacey.flf",
	"Stampate.flf",
	"Stampatello.flf",
	"Standard.flf",
	"Star Strips.flf",
	"Star Wars.flf",
	"Stellar.flf",
	"Stforek.flf",
	"Stick Letters.flf",
	"Stop.flf",
	"Straight.flf",
	"Stronger Than All.flf",
	"Sub-Zero.flf",
	"Swamp Land.flf",
	"Swan.flf",
	"Sweet.flf",
	"THIS.flf",
	"Tanja.flf",
	"Tengwar.flf",
	"Term.flf",
	"Test1.flf",
	"The Edge.flf",
	"Thick.flf",
	"Thin.flf",
	"Thorned.flf",
	"Three Point.flf",
	"Ticks Slant.flf",
	"Ticks.flf",
	"Tiles.flf",
	"Tinker-Toy.flf",
	"Tombstone.flf",
	"Train.flf",
	"Trek.flf",
	"Tsalagi.flf",
	"Tubular.flf",
	"Twisted.flf",
	"Two Point.flf",
	"USA Flag.flf",
	"Univers.flf",
	"Varsity.flf",
	"Wavy.flf",
	"Weird.flf",
	"Wet Letter.flf",
	"Whimsy.flf",
	"Wow.flf",
	"alligator3.flf",
	"amc3line.flf",
	"amc3liv1.flf",
	"amcaaa01.flf",
	"amcneko.flf",
	"amcrazo2.flf",
	"amcrazor.flf",
	"amcslash.flf",
	"amcslder.flf",
	"amcthin.flf",
	"amctubes.flf",
	"amcun1.flf",
	"bigchief.flf",
	"broadway_kb.flf",
	"calgphy2.flf",
	"cosmic.flf",
	"dosrebel.flf",
	"dotmatrix.flf",
	"doubleshorts.flf",
	"drpepper.flf",
	"eftichess.flf",
	"eftifont.flf",
	"eftipiti.flf",
	"eftirobot.flf",
	"eftitalic.flf",
	"eftiwall.flf",
	"eftiwater.flf",
	"fire_font-k.flf",
	"fire_font-s.flf",
	"fourtops.flf",
	"henry3d.flf",
	"horizontalleft.flf",
	"horizontalright.flf",
	"larry3d.flf",
	"lildevil.flf",
	"ntgreek.flf",
	"oldbanner.flf",
	"peaksslant.flf",
	"red_phoenix.flf",
	"rev.flf",
	"rowancap.flf",
	"s-relief.flf",
	"santaclara.flf",
	"sblood.flf",
	"slscript.flf",
	"smisome1.flf",
	"smkeyboard.flf",
	"smpoison.flf",
	"smscript.flf",
	"smshadow.flf",
	"smslant.flf",
	"smtengwar.flf",
	"starwars.flf",
	"threepoint.flf",
	"ticksslant.flf",
	"twopoint.flf",
	"usaflag.flf",
}
