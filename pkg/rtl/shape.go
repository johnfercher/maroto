package rtl

import "unicode"

// joiningForm indexes the contextual presentation forms of an Arabic letter.
//
// A letter is rendered differently depending on whether it connects to the
// letter before it, after it, both or neither. Fonts expose those variants
// through the Arabic Presentation Forms-B block (U+FE70..U+FEFF).
type joiningForm struct {
	isolated rune
	final    rune
	initial  rune
	medial   rune
	// dual reports whether the letter joins to the following letter as well as
	// to the preceding one. Letters that are not dual joining only ever take
	// the isolated or the final form.
	dual bool
}

// canJoinBackward reports whether the letter connects to the preceding letter,
// which is true for every Arabic letter that owns a final form. Hamza
// (U+0621) is the only non joining letter: it has no final form.
func (f joiningForm) canJoinBackward() bool { return f.final != 0 }

// canJoinForward reports whether the letter connects to the following letter.
// Only dual joining letters do.
func (f joiningForm) canJoinForward() bool { return f.dual }

// resolve returns the presentation form to use given the joining state of the
// neighbouring letters.
func (f joiningForm) resolve(prevJoins, nextJoins bool) rune {
	if !f.dual {
		if prevJoins && f.canJoinBackward() {
			return f.final
		}
		return f.isolated
	}

	switch {
	case prevJoins && nextJoins:
		return f.medial
	case prevJoins:
		return f.final
	case nextJoins:
		return f.initial
	default:
		return f.isolated
	}
}

// arabicForms maps the Arabic letters U+0621..U+063A and U+0641..U+064A to
// their contextual forms. Letters flagged dual are the dual joining ones; the
// remaining entries (the alef variants, waw, dal, thal, reh, zain, taa marbuta
// and alef maksura) are right joining and only own an isolated and a final
// form. Hamza owns neither, so it never connects.
var arabicForms = map[rune]joiningForm{
	0x0621: {isolated: 0xFE80},                                                             // hamza
	0x0622: {isolated: 0xFE81, final: 0xFE82},                                              // alef with madda above
	0x0623: {isolated: 0xFE83, final: 0xFE84},                                              // alef with hamza above
	0x0624: {isolated: 0xFE85, final: 0xFE86},                                              // waw with hamza above
	0x0625: {isolated: 0xFE87, final: 0xFE88},                                              // alef with hamza below
	0x0626: {isolated: 0xFE89, final: 0xFE8A, initial: 0xFE8B, medial: 0xFE8C, dual: true}, // yeh with hamza above
	0x0627: {isolated: 0xFE8D, final: 0xFE8E},                                              // alef
	0x0628: {isolated: 0xFE8F, final: 0xFE90, initial: 0xFE91, medial: 0xFE92, dual: true}, // beh
	0x0629: {isolated: 0xFE93, final: 0xFE94},                                              // taa marbuta
	0x062A: {isolated: 0xFE95, final: 0xFE96, initial: 0xFE97, medial: 0xFE98, dual: true}, // taa
	0x062B: {isolated: 0xFE99, final: 0xFE9A, initial: 0xFE9B, medial: 0xFE9C, dual: true}, // theh
	0x062C: {isolated: 0xFE9D, final: 0xFE9E, initial: 0xFE9F, medial: 0xFEA0, dual: true}, // jeem
	0x062D: {isolated: 0xFEA1, final: 0xFEA2, initial: 0xFEA3, medial: 0xFEA4, dual: true}, // hah
	0x062E: {isolated: 0xFEA5, final: 0xFEA6, initial: 0xFEA7, medial: 0xFEA8, dual: true}, // khah
	0x062F: {isolated: 0xFEA9, final: 0xFEAA},                                              // dal
	0x0630: {isolated: 0xFEAB, final: 0xFEAC},                                              // thal
	0x0631: {isolated: 0xFEAD, final: 0xFEAE},                                              // reh
	0x0632: {isolated: 0xFEAF, final: 0xFEB0},                                              // zain
	0x0633: {isolated: 0xFEB1, final: 0xFEB2, initial: 0xFEB3, medial: 0xFEB4, dual: true}, // seen
	0x0634: {isolated: 0xFEB5, final: 0xFEB6, initial: 0xFEB7, medial: 0xFEB8, dual: true}, // sheen
	0x0635: {isolated: 0xFEB9, final: 0xFEBA, initial: 0xFEBB, medial: 0xFEBC, dual: true}, // sad
	0x0636: {isolated: 0xFEBD, final: 0xFEBE, initial: 0xFEBF, medial: 0xFEC0, dual: true}, // dad
	0x0637: {isolated: 0xFEC1, final: 0xFEC2, initial: 0xFEC3, medial: 0xFEC4, dual: true}, // tah
	0x0638: {isolated: 0xFEC5, final: 0xFEC6, initial: 0xFEC7, medial: 0xFEC8, dual: true}, // zah
	0x0639: {isolated: 0xFEC9, final: 0xFECA, initial: 0xFECB, medial: 0xFECC, dual: true}, // ain
	0x063A: {isolated: 0xFECD, final: 0xFECE, initial: 0xFECF, medial: 0xFED0, dual: true}, // ghain
	0x0641: {isolated: 0xFED1, final: 0xFED2, initial: 0xFED3, medial: 0xFED4, dual: true}, // feh
	0x0642: {isolated: 0xFED5, final: 0xFED6, initial: 0xFED7, medial: 0xFED8, dual: true}, // qaf
	0x0643: {isolated: 0xFED9, final: 0xFEDA, initial: 0xFEDB, medial: 0xFEDC, dual: true}, // kaf
	0x0644: {isolated: 0xFEDD, final: 0xFEDE, initial: 0xFEDF, medial: 0xFEE0, dual: true}, // lam
	0x0645: {isolated: 0xFEE1, final: 0xFEE2, initial: 0xFEE3, medial: 0xFEE4, dual: true}, // meem
	0x0646: {isolated: 0xFEE5, final: 0xFEE6, initial: 0xFEE7, medial: 0xFEE8, dual: true}, // noon
	0x0647: {isolated: 0xFEE9, final: 0xFEEA, initial: 0xFEEB, medial: 0xFEEC, dual: true}, // heh
	0x0648: {isolated: 0xFEED, final: 0xFEEE},                                              // waw
	0x0649: {isolated: 0xFEEF, final: 0xFEF0},                                              // alef maksura
	0x064A: {isolated: 0xFEF1, final: 0xFEF2, initial: 0xFEF3, medial: 0xFEF4, dual: true}, // yeh
}

// lam is the letter that forms the mandatory ligatures with the alef variants.
const lam = 0x0644

// lamAlefLigatures maps the alef that follows a lam to the ligature the pair
// must contract into. These ligatures are mandatory in Arabic typography: a
// lam followed by an alef is never rendered as two separate glyphs. The
// resulting ligature is right joining, so it takes the final form when the lam
// connects to the preceding letter and the isolated form otherwise.
var lamAlefLigatures = map[rune]joiningForm{
	0x0622: {isolated: 0xFEF5, final: 0xFEF6}, // lam + alef with madda above
	0x0623: {isolated: 0xFEF7, final: 0xFEF8}, // lam + alef with hamza above
	0x0625: {isolated: 0xFEF9, final: 0xFEFA}, // lam + alef with hamza below
	0x0627: {isolated: 0xFEFB, final: 0xFEFC}, // lam + alef
}

// isCombining reports whether r is a non spacing mark that must stay glued to
// the base character preceding it, both when shaping and when reversing.
//
// This covers the Arabic marks U+064B..U+0652 (tanween, fatha, damma, kasra,
// shadda and sukun), which are transparent for joining: they never break the
// connection between the letters surrounding them.
func isCombining(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// cluster is a base character together with the combining marks that follow
// it. Shaping and reordering both operate on clusters so that a mark never
// gets detached from the character it belongs to.
type cluster struct {
	base  rune
	marks []rune
	// form is set when base is an Arabic letter that must be replaced by a
	// contextual presentation form. It is nil for every other character, which
	// is therefore emitted unchanged.
	form *joiningForm
}

// splitClusters breaks text into clusters, attaching every combining mark to
// the base character that precedes it.
func splitClusters(text string) []cluster {
	clusters := make([]cluster, 0, len(text))

	for _, r := range text {
		if isCombining(r) && len(clusters) > 0 {
			last := len(clusters) - 1
			clusters[last].marks = append(clusters[last].marks, r)
			continue
		}
		clusters = append(clusters, cluster{base: r})
	}

	return clusters
}

// contractLigatures replaces every lam followed by an alef variant with the
// corresponding mandatory ligature. Marks sitting between the lam and the alef
// do not prevent the contraction, they are simply carried over to the ligature
// in their original order, which is what real typography does.
func contractLigatures(clusters []cluster) []cluster {
	out := make([]cluster, 0, len(clusters))

	for i := 0; i < len(clusters); i++ {
		current := clusters[i]

		if current.base == lam && i+1 < len(clusters) {
			if ligature, ok := lamAlefLigatures[clusters[i+1].base]; ok {
				alef := clusters[i+1]
				merged := cluster{base: current.base, form: &ligature}
				merged.marks = append(merged.marks, current.marks...)
				merged.marks = append(merged.marks, alef.marks...)
				out = append(out, merged)
				i++

				continue
			}
		}

		if form, ok := arabicForms[current.base]; ok {
			current.form = &form
		}

		out = append(out, current)
	}

	return out
}

// shape replaces the Arabic letters of text by their contextual presentation
// forms, leaving every other character untouched. The returned string is still
// in logical order; reordering is a separate step.
func shape(text string) string {
	clusters := contractLigatures(splitClusters(text))

	var builder []rune
	for i, current := range clusters {
		builder = append(builder, shapeCluster(clusters, i, current))
		builder = append(builder, current.marks...)
	}

	return string(builder)
}

// shapeCluster resolves the presentation form of the cluster at index i.
func shapeCluster(clusters []cluster, i int, current cluster) rune {
	if current.form == nil {
		return current.base
	}

	return current.form.resolve(joinsForward(clusters, i-1), joinsBackward(clusters, i+1))
}

// joinsForward reports whether the cluster at index i connects to the cluster
// that follows it. Anything that is not an Arabic letter breaks the joining.
func joinsForward(clusters []cluster, i int) bool {
	if i < 0 || i >= len(clusters) || clusters[i].form == nil {
		return false
	}

	return clusters[i].form.canJoinForward()
}

// joinsBackward reports whether the cluster at index i connects to the cluster
// that precedes it.
func joinsBackward(clusters []cluster, i int) bool {
	if i < 0 || i >= len(clusters) || clusters[i].form == nil {
		return false
	}

	return clusters[i].form.canJoinBackward()
}
