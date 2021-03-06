// Code generated by gocode.Generate; DO NOT EDIT.

package filetypes

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalFileInfo = cuegenMake("FileInfo", &FileInfo{})

// Validate validates x.
func (x *FileInfo) Validate() error {
	return cuegenCodec.Validate(cuegenvalFileInfo, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 1132 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xdcWO\x8b\xdc6\x14\xb7&[\xa8DR\xda\x0fPP}\b\xe9@\xe7Z\x18\b\xb9\xa4\x81\\J\xe9u\t\x8b\u05a3\xf1\xba\xb1%c\xcbe\x97\xee\x1c\u06a6i?p\u03d9\xf2\xf4$\u02d2=Iw!P\xaa\u02cc\u007fO\xef\xff{\xd2\u04e3\xe3\x9f+\xb2:\xfe\x95\x91\xe3oY\xf6\xed\xf1\xd7\a\x84<\xacTo\x84*\xe4sa\x04\xe0\xe4\x019\xfbQkCV\x199\xfbA\x98+\xf20#\x9f\xbc\xa8j\u0653\xe3\xdb,\u02fe<\xfe\xb1\"\xe4\xb3\xf3W\xc5 7\xfb\xaav\x9co3r|\x93eO\x8e\xbf? \xe4\u04c0\xbf\xc9\u020a\x9c}/\x1a\t\x82\xce,\u0232,{\xf7\xc5#\xb0\x84\x90\x15!\xd4\u0734\xb2\xdf\x14\x83$\xef>\xff\xbb\x15\xc5kQJ~9T\xf5\x8e1P\u0377[\xfe\v\xa3 U\x89Fn\xb9[\xbd\xe9*U2*U\xa1w\x95*G\xc2w\x0e`\xb4RFvm'\x8d0\x95V\u03f6\xfce\x040\xba\xd7]\xf3ld\u473f\xd0]\u00e8\x11e\xff\xccj\xa5\xe7\xa8\xe6\xd5v\xd4w`\a\xe6U\x80m\xb8\xbez\x9a\xe7,\x16\x0fD\xc7\x04b\xc3\u0789\xf5\xa0\xc9*2\xf2\u06a0\xc6\xe0O\x0e`\u03a85\x13\x99\xf3\x9d0\"\a#(\xfcC\x0e$OH\xc5 g\xb2\x8aA\"\xb1/\xaed\x13s\"\x84\xe4\x9fz\xadf\xcc\x00\x06r\xbdH\xafq\u00cdh\xe6t\x00\x91\\\ua508+/\xf4\x0e\fLr\xb6\xe59\x80c\x98(\xad\x85\x95Xj\xc0\x0fVf\xdbi3\x13\x9b[\xd4)\xedD{\x15yl\x11\x1f\xc72\tc\u98a8w\xb30\xde\xcbX_9\xce\\1\u032d\u0175\xc6@\xf3[~\xb1$\x1c\x18C\x16\xa6y\xbc\xa3\xa0\xc0\x8e\xe2t+\x95h\xab{\xc9r\xbc96\xc6s\xb9\x17Cm\xa0\xd6m\xef>\x8e[w\x9d\u007f\x03\x82\\@\x0e\xb6\xbf_\xaa\xbdv=\x8e%\xed\xd7\xdat\x83\xe4\xb7|/\xea^2\xda\u027d\xec\xa4*d\xbf\x9d\x13\x8b\x9b\xa2F\xc2\x02\xe7N\xee+U\x81\xbd\xb0\xe3R\xeb\x1a\\\x86oQ#\vb\x85V\xbd\xe9D\xa5L\xd8\xf7Z\xca\xd69\xd5o\x1dV\xa9B7m-\x8d=\x8c\x1c\u05b4\xba3\xde\x02\xc4z\xd3I\xd1x\xa3\x10\xdb\xe9\xa2\x0f.\"&\x8c\xe9\xaa\xcb\xc1\xa0\x03\x88Ad\u06015z'\xb1\x98*\xd5\x0e\ue118\x04\xd9VY\xc8\xd8\u06b6\xb9K\x15\xddl6Xt4\x893\x8d\xccH\xa2E#{f\xc4 \u052b\xf5\xc5>\x9e:\x94\xae\xa1\x97\xfa\r\u05987\xe7\xe0\xf9\xae\x8dT=f\xc3n\xcf7\xb6\xc2<sZbk,\xfcH\f\x16\x16\xb0\u06a3\u57acw\xe4\xb4%\x8e\xec\xf2\x1a\xd2\xfd\xe1tL[\xe7\xdf\xe6\xc3E|\xbd\x98\x8f\x19\xf1.\xf9\x80\xfe:\xe5\x91\x1eZ\xe3\v\xec?h\x9e\xfcY\xd4\x1f\xa5\xfa?\x9e;\xfbJ\x89\xfa\x94?;\xb9\xff?43\x9c\xf91\xeb\xc8\xe8O?\xa7l\xf4<\f\x18\xd3\xf3\x11C\xccom\xc8Y\xec\x937w\xea\x8a\xc7B\xe5LT\x84\x19$2\x022\xe9\xc4#\x9b\x9d\x02\x12\xd3\x12\xc6h\u007f\xa2&\xccl\x89\xaf'\xb6\xebE\xabNm\u007f\xffTC\xa7I\x99p\x85\xa1eY\u0244\x81Y\xc4n>\x87\x19\xfd\xe98\xea\xfa\n\xf3B\xf3|\xcb/\xfc\xc7|~\x1c/l?H\xf2[\x9e\xdb\xea\xb7\xff\xfc\xac\x95\xdc\xc3i1\xc67\xf2\x93\x88\xfc5\u007f\x9c\"\x8c&\xf7u*/\xbe\xb9Sj|\x87\u03e8\xd1m\x9eR\xe3{=9#\xc6L\xd8\x00,\x85\u0245f\xe6\xf2\xb2\xe1'\u04c7Zf\xa3mH\x06\xc6\x1d2\x00#-\xfe\xda\x17B2g\xf9\x8e\x8b\xb2\xb3\x9c\x15\x8f\u0191\u007f\xbf\xe1q\xa4\x97#\x9c\xc6.\x99\xc9\xed\xf3j,#?\xeb\u0151I{9}\x1e\x85w\x1a\x0e\xf8\xf1<9:6\x9d#\x17c\xb0\x18\x82E\xaf\xd2\xeevo\xc7d\xfc\x81\xder\ub085Q\x88sxF\xd8/\x16\xa6\x1c\x87\xc2\x17\x1b\a\x98\x14\x85[#\xa0\x05$\x17\xc5\xda\u0363X\xbb\xb7\xde9}1\xac\x96asm&\x92\xe1\xf8\x03\xb4\xd4\xce\x05\x8b\x96\x1a0<\xa6\xbc:\xfb\x05\aNU\xcbp\xba\xa4\x8f\x83|\xaf\xf5\u01bd4'/^\xffP9\xb0xF\xbb\xfb\xc9\x15\x1e^'\x9a\xe9\xf4\xb3*\x9e\xe4N\xb0\x9fxF}\x80\x97e\xd9?\x01\x00\x00\xff\xffjB\xe5G\x9d\x11\x00\x00")
