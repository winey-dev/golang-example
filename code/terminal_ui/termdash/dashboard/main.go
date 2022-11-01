package main

import (
	"context"
	"dashboard/pkg/dashboard"
	"fmt"
	"image"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
)

var rootID string = "root"

/*
ColorMaroon
ColorGreen
ColorOlive
ColorNavy
ColorPurple
ColorTeal
ColorSilver
ColorGray
ColorRed
ColorLime
ColorYellow
ColorBlue
ColorFuchsia
ColorAqua
ColorWhite
*/

var colormap = map[int]cell.Color{
	0:  cell.ColorMaroon,
	1:  cell.ColorOlive,
	2:  cell.ColorNavy,
	3:  cell.ColorPurple,
	5:  cell.ColorTeal,
	6:  cell.ColorSilver,
	7:  cell.ColorGray,
	9:  cell.ColorRed,
	10: cell.ColorLime,
	11: cell.ColorYellow,
	12: cell.ColorBlue,
	13: cell.ColorFuchsia,
	14: cell.ColorAqua,
	15: cell.ColorWhite,
}

const redrawInterval = 250 * time.Millisecond

func main() {
	t, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256))
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	// containers new
	c, err := container.New(t, container.ID(rootID))
	if err != nil {
		panic(err)
	}

	title, err := newDisplay(ctx, t)
	if err != nil {
		panic(err)
	}

	charts := newLineChart(ctx)

	chartInfo, err := newText()
	if err != nil {
		panic(err)
	}

	gridOpts, err := gridLayout(title, charts, chartInfo)
	if err != nil {
		panic(err)
	}

	if err := c.Update(rootID, gridOpts...); err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(redrawInterval)); err != nil {
		panic(err)
	}

}

func newText() ([]*text.Text, error) {
	items := []*text.Text{}

	for _, chart := range dashboard.Data.Charts {
		t, err := text.New(text.WrapAtRunes())
		if err != nil {
			return nil, err
		}
		for i, col := range chart.Cols {
			t.Write(col+"\n", text.WriteCellOpts(cell.FgColor(colormap[i])))
		}
		items = append(items, t)
	}

	return items, nil
}

/*Dis Play Func Start */
func newDisplay(ctx context.Context, t terminalapi.Terminal) (*segmentdisplay.SegmentDisplay, error) {
	sd, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}

	colors := []cell.Color{
		cell.ColorNumber(33),
		cell.ColorRed,
		cell.ColorYellow,
		cell.ColorNumber(33),
		cell.ColorGreen,
		cell.ColorRed,
		cell.ColorGreen,
		cell.ColorRed,
	}

	text := dashboard.Data.Name
	step := 0

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		capacity := 0
		termSize := t.Size()
		for {
			select {
			case <-ticker.C:
				if capacity == 0 {
					capacity = sd.Capacity()
				}
				if t.Size().Eq(image.ZP) || !t.Size().Eq(termSize) {
					termSize = t.Size()
					capacity = sd.Capacity()
				}

				state := textState(text, capacity, step)
				var chunks []*segmentdisplay.TextChunk
				for i := 0; i < capacity; i++ {
					if i >= len(state) {
						break
					}

					color := colors[i%len(colors)]
					chunks = append(chunks, segmentdisplay.NewChunk(
						string(state[i]),
						segmentdisplay.WriteCellOpts(cell.FgColor(color)),
					))
				}
				if len(chunks) == 0 {
					continue
				}
				if err := sd.Write(chunks); err != nil {
					panic(err)
				}
				step++
			case <-ctx.Done():
				return
			}
		}
	}()
	return sd, nil

}

func rotateRunes(inputs []rune, step int) []rune {
	return append(inputs[step:], inputs[:step]...)
}

func textState(text string, capacity, step int) []rune {
	if capacity == 0 {
		return nil
	}

	var state []rune
	for i := 0; i < capacity; i++ {
		state = append(state, ' ')
	}
	state = append(state, []rune(text)...)
	step = step % len(state)
	return rotateRunes(state, step)
}

/*Dis Play Func End */
func newLineChart(ctx context.Context) []*linechart.LineChart {
	items := []*linechart.LineChart{}
	for _, chart := range dashboard.Data.Charts {
		items = append(items, newLineChartInit(ctx, chart))
	}
	return items
}

func newLineChartInit(ctx context.Context, chart *dashboard.Chart) *linechart.LineChart {
	lc, _ := linechart.New(
		linechart.AxesCellOpts(cell.FgColor(cell.ColorRed)),
		linechart.YLabelCellOpts(cell.FgColor(cell.ColorGreen)),
		linechart.XLabelCellOpts(cell.FgColor(cell.ColorGreen)),
	)

	chart.InitValue(dashboard.Data.Interval, dashboard.Data.Range)
	xLabel := chart.MakeLabel()
	for i, col := range chart.Cols {
		lc.Series(col, chart.Values[i],
			linechart.SeriesCellOpts(cell.FgColor(colormap[i])),
			linechart.SeriesXLabels(xLabel),
		)
	}

	go func() {
		ticker := time.NewTicker(dashboard.Data.Interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				chart.RotateFloats()
				xLabel := chart.MakeLabel()
				for i, col := range chart.Cols {
					lc.Series(col, chart.Values[i],
						linechart.SeriesCellOpts(cell.FgColor(colormap[i])),
						linechart.SeriesXLabels(xLabel),
					)
				}
			case <-ctx.Done():
				return
			}
		}

	}()
	return lc
}

func gridLayout(sd *segmentdisplay.SegmentDisplay, lcs []*linechart.LineChart, chartInfo []*text.Text) ([]container.Option, error) {
	fram := []grid.Element{
		grid.RowHeightPerc(20,
			grid.Widget(sd,
				container.Border(linestyle.Light),
				container.BorderTitle("Press Esc to quit"),
			),
		),
	}

	// 5개 이상부터 right 활성화
	left := []grid.Element{}
	right := []grid.Element{}
	var leftHeight int
	var rightHeight int
	var rightIndex, leftIndex int

	if len(lcs) < 5 {
		leftIndex = len(lcs)
		leftHeight = 100 / len(lcs)
	} else {
		rightIndex = len(lcs) / 2
		leftIndex = len(lcs) - rightIndex
		leftHeight = 100 / leftIndex
		rightHeight = 100 / rightIndex
	}

	for i, lc := range lcs {
		if i < leftIndex {
			left = append(left, grid.RowHeightPerc(leftHeight,
				grid.ColWidthPerc(80,
					grid.Widget(lc,
						container.Border(linestyle.Light),
						container.BorderTitle(dashboard.Data.Charts[i].Name),
					),
				),
				grid.ColWidthPerc(20,
					grid.Widget(chartInfo[i],
						container.Border(linestyle.Light),
						container.BorderTitle(fmt.Sprintf("%s line information", dashboard.Data.Charts[i].Name)),
					),
				),
			))
		} else {
			right = append(right, grid.RowHeightPerc(rightHeight,
				grid.ColWidthPerc(80,
					grid.Widget(lc,
						container.Border(linestyle.Light),
						container.BorderTitle(dashboard.Data.Charts[i].Name),
					),
				),
				grid.ColWidthPerc(20,
					grid.Widget(chartInfo[i],
						container.Border(linestyle.Light),
						container.BorderTitle(fmt.Sprintf("%s line information", dashboard.Data.Charts[i].Name)),
					),
				),
			))
			/*
				right = append(right, grid.RowHeightPerc(rightHeight,
					grid.Widget(lc,
						container.Border(linestyle.Light),
						container.BorderTitle(dashboard.Data.Charts[i].Name),
					),
				))
			*/
		}
	}

	if len(right) != 0 {
		fram = append(fram,
			grid.RowHeightPerc(80,
				grid.ColWidthPerc(50, left...),
				grid.ColWidthPerc(50, right...),
			),
		)
	} else {
		fram = append(fram,
			grid.RowHeightPerc(80,
				left...,
			),
		)

	}

	builder := grid.New()
	builder.Add(
		fram...,
	)

	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}

	return gridOpts, nil

}
