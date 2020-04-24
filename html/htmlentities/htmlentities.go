/**
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 **/

package htmlentities

type Entity struct {
	Codepoints []int64
	Characters string
}

var List = map[string]Entity{
	"&ldca;": Entity{
		Codepoints: []int64{10550},
		Characters: "\u2936",
	},
	"&conint;": Entity{
		Codepoints: []int64{8751},
		Characters: "\u222f",
	},
	"&equiv;": Entity{
		Codepoints: []int64{8801},
		Characters: "\u2261",
	},
	"&aogon;": Entity{
		Codepoints: []int64{260},
		Characters: "\u0104",
	},
	"&ogt;": Entity{
		Codepoints: []int64{10689},
		Characters: "\u29c1",
	},
	"&uharl;": Entity{
		Codepoints: []int64{8639},
		Characters: "\u21bf",
	},
	"&gfr;": Entity{
		Codepoints: []int64{120074},
		Characters: "\U0001d50a",
	},
	"&eddot;": Entity{
		Codepoints: []int64{10871},
		Characters: "\u2a77",
	},
	"&toea;": Entity{
		Codepoints: []int64{10536},
		Characters: "\u2928",
	},
	"&cularr;": Entity{
		Codepoints: []int64{8630},
		Characters: "\u21b6",
	},
	"&bsolb;": Entity{
		Codepoints: []int64{10693},
		Characters: "\u29c5",
	},
	"&sime;": Entity{
		Codepoints: []int64{8771},
		Characters: "\u2243",
	},
	"&upsi;": Entity{
		Codepoints: []int64{965},
		Characters: "\u03c5",
	},
	"&awconint;": Entity{
		Codepoints: []int64{8755},
		Characters: "\u2233",
	},
	"&boxvl;": Entity{
		Codepoints: []int64{9508},
		Characters: "\u2524",
	},
	"&ecaron;": Entity{
		Codepoints: []int64{283},
		Characters: "\u011b",
	},
	"&awint;": Entity{
		Codepoints: []int64{10769},
		Characters: "\u2a11",
	},
	"&cylcty;": Entity{
		Codepoints: []int64{9005},
		Characters: "\u232d",
	},
	"&ycy;": Entity{
		Codepoints: []int64{1067},
		Characters: "\u042b",
	},
	"&nsqsube;": Entity{
		Codepoints: []int64{8930},
		Characters: "\u22e2",
	},
	"&udblac;": Entity{
		Codepoints: []int64{368},
		Characters: "\u0170",
	},
	"&efr;": Entity{
		Codepoints: []int64{120072},
		Characters: "\U0001d508",
	},
	"&commat;": Entity{
		Codepoints: []int64{64},
		Characters: "@",
	},
	"&ubrcy;": Entity{
		Codepoints: []int64{1038},
		Characters: "\u040e",
	},
	"&ac;": Entity{
		Codepoints: []int64{8766},
		Characters: "\u223e",
	},
	"&jmath;": Entity{
		Codepoints: []int64{567},
		Characters: "\u0237",
	},
	"&bopf;": Entity{
		Codepoints: []int64{120147},
		Characters: "\U0001d553",
	},
	"&larrb;": Entity{
		Codepoints: []int64{8676},
		Characters: "\u21e4",
	},
	"&nvgt;": Entity{
		Codepoints: []int64{62, 8402},
		Characters: ">\u20d2",
	},
	"&thetav;": Entity{
		Codepoints: []int64{977},
		Characters: "\u03d1",
	},
	"&angmsd;": Entity{
		Codepoints: []int64{8737},
		Characters: "\u2221",
	},
	"&aring;": Entity{
		Codepoints: []int64{229},
		Characters: "\u00e5",
	},
	"&lsim;": Entity{
		Codepoints: []int64{8818},
		Characters: "\u2272",
	},
	"&nvrarr;": Entity{
		Codepoints: []int64{10499},
		Characters: "\u2903",
	},
	"&otimesas;": Entity{
		Codepoints: []int64{10806},
		Characters: "\u2a36",
	},
	"&lacute;": Entity{
		Codepoints: []int64{314},
		Characters: "\u013a",
	},
	"&nape;": Entity{
		Codepoints: []int64{10864, 824},
		Characters: "\u2a70\u0338",
	},
	"&rlarr;": Entity{
		Codepoints: []int64{8644},
		Characters: "\u21c4",
	},
	"&dscy;": Entity{
		Codepoints: []int64{1029},
		Characters: "\u0405",
	},
	"&frac15;": Entity{
		Codepoints: []int64{8533},
		Characters: "\u2155",
	},
	"&comp;": Entity{
		Codepoints: []int64{8705},
		Characters: "\u2201",
	},
	"&icy;": Entity{
		Codepoints: []int64{1080},
		Characters: "\u0438",
	},
	"&nwarhk;": Entity{
		Codepoints: []int64{10531},
		Characters: "\u2923",
	},
	"&prurel;": Entity{
		Codepoints: []int64{8880},
		Characters: "\u22b0",
	},
	"&rarr;": Entity{
		Codepoints: []int64{8608},
		Characters: "\u21a0",
	},
	"&nscr;": Entity{
		Codepoints: []int64{120003},
		Characters: "\U0001d4c3",
	},
	"&cup;": Entity{
		Codepoints: []int64{8746},
		Characters: "\u222a",
	},
	"&imath;": Entity{
		Codepoints: []int64{305},
		Characters: "\u0131",
	},
	"&acute;": Entity{
		Codepoints: []int64{180},
		Characters: "\u00b4",
	},
	"&orv;": Entity{
		Codepoints: []int64{10843},
		Characters: "\u2a5b",
	},
	"&rlm;": Entity{
		Codepoints: []int64{8207},
		Characters: "\u200f",
	},
	"&sigmav;": Entity{
		Codepoints: []int64{962},
		Characters: "\u03c2",
	},
	"&hbar;": Entity{
		Codepoints: []int64{8463},
		Characters: "\u210f",
	},
	"&thorn;": Entity{
		Codepoints: []int64{254},
		Characters: "\u00fe",
	},
	"&uharr;": Entity{
		Codepoints: []int64{8638},
		Characters: "\u21be",
	},
	"&bsolhsub;": Entity{
		Codepoints: []int64{10184},
		Characters: "\u27c8",
	},
	"&triplus;": Entity{
		Codepoints: []int64{10809},
		Characters: "\u2a39",
	},
	"&sfr;": Entity{
		Codepoints: []int64{120086},
		Characters: "\U0001d516",
	},
	"&mcy;": Entity{
		Codepoints: []int64{1084},
		Characters: "\u043c",
	},
	"&in;": Entity{
		Codepoints: []int64{8712},
		Characters: "\u2208",
	},
	"&origof;": Entity{
		Codepoints: []int64{8886},
		Characters: "\u22b6",
	},
	"&elsdot;": Entity{
		Codepoints: []int64{10903},
		Characters: "\u2a97",
	},
	"&dopf;": Entity{
		Codepoints: []int64{120123},
		Characters: "\U0001d53b",
	},
	"&xdtri;": Entity{
		Codepoints: []int64{9661},
		Characters: "\u25bd",
	},
	"&plusb;": Entity{
		Codepoints: []int64{8862},
		Characters: "\u229e",
	},
	"&block;": Entity{
		Codepoints: []int64{9608},
		Characters: "\u2588",
	},
	"&times;": Entity{
		Codepoints: []int64{215},
		Characters: "\u00d7",
	},
	"&robrk;": Entity{
		Codepoints: []int64{10215},
		Characters: "\u27e7",
	},
	"&sqsup;": Entity{
		Codepoints: []int64{8848},
		Characters: "\u2290",
	},
	"&tshcy;": Entity{
		Codepoints: []int64{1115},
		Characters: "\u045b",
	},
	"&gle;": Entity{
		Codepoints: []int64{10898},
		Characters: "\u2a92",
	},
	"&sqcaps;": Entity{
		Codepoints: []int64{8851, 65024},
		Characters: "\u2293\ufe00",
	},
	"&imacr;": Entity{
		Codepoints: []int64{299},
		Characters: "\u012b",
	},
	"&drcorn;": Entity{
		Codepoints: []int64{8991},
		Characters: "\u231f",
	},
	"&dd;": Entity{
		Codepoints: []int64{8517},
		Characters: "\u2145",
	},
	"&rightdownvectorbar;": Entity{
		Codepoints: []int64{10581},
		Characters: "\u2955",
	},
	"&lharu;": Entity{
		Codepoints: []int64{8636},
		Characters: "\u21bc",
	},
	"&fopf;": Entity{
		Codepoints: []int64{120125},
		Characters: "\U0001d53d",
	},
	"&odsold;": Entity{
		Codepoints: []int64{10684},
		Characters: "\u29bc",
	},
	"&bumpe;": Entity{
		Codepoints: []int64{8783},
		Characters: "\u224f",
	},
	"&rbrksld;": Entity{
		Codepoints: []int64{10638},
		Characters: "\u298e",
	},
	"&alpha;": Entity{
		Codepoints: []int64{945},
		Characters: "\u03b1",
	},
	"&midcir;": Entity{
		Codepoints: []int64{10992},
		Characters: "\u2af0",
	},
	"&olcir;": Entity{
		Codepoints: []int64{10686},
		Characters: "\u29be",
	},
	"&minusdu;": Entity{
		Codepoints: []int64{10794},
		Characters: "\u2a2a",
	},
	"&supsub;": Entity{
		Codepoints: []int64{10964},
		Characters: "\u2ad4",
	},
	"&yicy;": Entity{
		Codepoints: []int64{1031},
		Characters: "\u0407",
	},
	"&barv;": Entity{
		Codepoints: []int64{10983},
		Characters: "\u2ae7",
	},
	"&vnsub;": Entity{
		Codepoints: []int64{8834, 8402},
		Characters: "\u2282\u20d2",
	},
	"&tau;": Entity{
		Codepoints: []int64{932},
		Characters: "\u03a4",
	},
	"&smeparsl;": Entity{
		Codepoints: []int64{10724},
		Characters: "\u29e4",
	},
	"&sdot;": Entity{
		Codepoints: []int64{8901},
		Characters: "\u22c5",
	},
	"&lfr;": Entity{
		Codepoints: []int64{120079},
		Characters: "\U0001d50f",
	},
	"&cuvee;": Entity{
		Codepoints: []int64{8910},
		Characters: "\u22ce",
	},
	"&urcrop;": Entity{
		Codepoints: []int64{8974},
		Characters: "\u230e",
	},
	"&ord;": Entity{
		Codepoints: []int64{10845},
		Characters: "\u2a5d",
	},
	"&zeta;": Entity{
		Codepoints: []int64{950},
		Characters: "\u03b6",
	},
	"&nvrtrie;": Entity{
		Codepoints: []int64{8885, 8402},
		Characters: "\u22b5\u20d2",
	},
	"&ldrushar;": Entity{
		Codepoints: []int64{10571},
		Characters: "\u294b",
	},
	"&nsub;": Entity{
		Codepoints: []int64{8836},
		Characters: "\u2284",
	},
	"&loang;": Entity{
		Codepoints: []int64{10220},
		Characters: "\u27ec",
	},
	"&daleth;": Entity{
		Codepoints: []int64{8504},
		Characters: "\u2138",
	},
	"&napid;": Entity{
		Codepoints: []int64{8779, 824},
		Characters: "\u224b\u0338",
	},
	"&zigrarr;": Entity{
		Codepoints: []int64{8669},
		Characters: "\u21dd",
	},
	"&sopf;": Entity{
		Codepoints: []int64{120138},
		Characters: "\U0001d54a",
	},
	"&veebar;": Entity{
		Codepoints: []int64{8891},
		Characters: "\u22bb",
	},
	"&odash;": Entity{
		Codepoints: []int64{8861},
		Characters: "\u229d",
	},
	"&breve;": Entity{
		Codepoints: []int64{728},
		Characters: "\u02d8",
	},
	"&amacr;": Entity{
		Codepoints: []int64{256},
		Characters: "\u0100",
	},
	"&elinters;": Entity{
		Codepoints: []int64{9191},
		Characters: "\u23e7",
	},
	"&rangd;": Entity{
		Codepoints: []int64{10642},
		Characters: "\u2992",
	},
	"&sqcup;": Entity{
		Codepoints: []int64{8852},
		Characters: "\u2294",
	},
	"&harr;": Entity{
		Codepoints: []int64{8596},
		Characters: "\u2194",
	},
	"&lrm;": Entity{
		Codepoints: []int64{8206},
		Characters: "\u200e",
	},
	"&lates;": Entity{
		Codepoints: []int64{10925, 65024},
		Characters: "\u2aad\ufe00",
	},
	"&grave;": Entity{
		Codepoints: []int64{96},
		Characters: "`",
	},
	"&smtes;": Entity{
		Codepoints: []int64{10924, 65024},
		Characters: "\u2aac\ufe00",
	},
	"&lnsim;": Entity{
		Codepoints: []int64{8934},
		Characters: "\u22e6",
	},
	"&nges;": Entity{
		Codepoints: []int64{10878, 824},
		Characters: "\u2a7e\u0338",
	},
	"&nrarrc;": Entity{
		Codepoints: []int64{10547, 824},
		Characters: "\u2933\u0338",
	},
	"&leg;": Entity{
		Codepoints: []int64{8922},
		Characters: "\u22da",
	},
	"&rcy;": Entity{
		Codepoints: []int64{1056},
		Characters: "\u0420",
	},
	"&rscr;": Entity{
		Codepoints: []int64{8475},
		Characters: "\u211b",
	},
	"&yen;": Entity{
		Codepoints: []int64{165},
		Characters: "\u00a5",
	},
	"&eng;": Entity{
		Codepoints: []int64{331},
		Characters: "\u014b",
	},
	"&rdca;": Entity{
		Codepoints: []int64{10551},
		Characters: "\u2937",
	},
	"&nmid;": Entity{
		Codepoints: []int64{8740},
		Characters: "\u2224",
	},
	"&vbar;": Entity{
		Codepoints: []int64{10987},
		Characters: "\u2aeb",
	},
	"&and;": Entity{
		Codepoints: []int64{10835},
		Characters: "\u2a53",
	},
	"&angmsdaa;": Entity{
		Codepoints: []int64{10664},
		Characters: "\u29a8",
	},
	"&ltrif;": Entity{
		Codepoints: []int64{9666},
		Characters: "\u25c2",
	},
	"&lparlt;": Entity{
		Codepoints: []int64{10643},
		Characters: "\u2993",
	},
	"&lcedil;": Entity{
		Codepoints: []int64{316},
		Characters: "\u013c",
	},
	"&nfr;": Entity{
		Codepoints: []int64{120107},
		Characters: "\U0001d52b",
	},
	"&colon;": Entity{
		Codepoints: []int64{8759},
		Characters: "\u2237",
	},
	"&gtlpar;": Entity{
		Codepoints: []int64{10645},
		Characters: "\u2995",
	},
	"&yfr;": Entity{
		Codepoints: []int64{120118},
		Characters: "\U0001d536",
	},
	"&otilde;": Entity{
		Codepoints: []int64{245},
		Characters: "\u00f5",
	},
	"&aleph;": Entity{
		Codepoints: []int64{8501},
		Characters: "\u2135",
	},
	"&caret;": Entity{
		Codepoints: []int64{8257},
		Characters: "\u2041",
	},
	"&ge;": Entity{
		Codepoints: []int64{8807},
		Characters: "\u2267",
	},
	"&sup3;": Entity{
		Codepoints: []int64{179},
		Characters: "\u00b3",
	},
	"&quest;": Entity{
		Codepoints: []int64{63},
		Characters: "?",
	},
	"&lg;": Entity{
		Codepoints: []int64{8822},
		Characters: "\u2276",
	},
	"&intlarhk;": Entity{
		Codepoints: []int64{10775},
		Characters: "\u2a17",
	},
	"&nvlt;": Entity{
		Codepoints: []int64{60, 8402},
		Characters: "<\u20d2",
	},
	"&barwed;": Entity{
		Codepoints: []int64{8965},
		Characters: "\u2305",
	},
	"&olt;": Entity{
		Codepoints: []int64{10688},
		Characters: "\u29c0",
	},
	"&gsime;": Entity{
		Codepoints: []int64{10894},
		Characters: "\u2a8e",
	},
	"&utrif;": Entity{
		Codepoints: []int64{9652},
		Characters: "\u25b4",
	},
	"&boxdl;": Entity{
		Codepoints: []int64{9558},
		Characters: "\u2556",
	},
	"&bepsi;": Entity{
		Codepoints: []int64{1014},
		Characters: "\u03f6",
	},
	"&lescc;": Entity{
		Codepoints: []int64{10920},
		Characters: "\u2aa8",
	},
	"&gsiml;": Entity{
		Codepoints: []int64{10896},
		Characters: "\u2a90",
	},
	"&vvdash;": Entity{
		Codepoints: []int64{8874},
		Characters: "\u22aa",
	},
	"&semi;": Entity{
		Codepoints: []int64{59},
		Characters: ";",
	},
	"&dash;": Entity{
		Codepoints: []int64{8208},
		Characters: "\u2010",
	},
	"&sim;": Entity{
		Codepoints: []int64{8764},
		Characters: "\u223c",
	},
	"&rx;": Entity{
		Codepoints: []int64{8478},
		Characters: "\u211e",
	},
	"&nlt;": Entity{
		Codepoints: []int64{8814},
		Characters: "\u226e",
	},
	"&scnsim;": Entity{
		Codepoints: []int64{8937},
		Characters: "\u22e9",
	},
	"&cacute;": Entity{
		Codepoints: []int64{263},
		Characters: "\u0107",
	},
	"&el;": Entity{
		Codepoints: []int64{10905},
		Characters: "\u2a99",
	},
	"&ddarr;": Entity{
		Codepoints: []int64{8650},
		Characters: "\u21ca",
	},
	"&bemptyv;": Entity{
		Codepoints: []int64{10672},
		Characters: "\u29b0",
	},
	"&beth;": Entity{
		Codepoints: []int64{8502},
		Characters: "\u2136",
	},
	"&epsiv;": Entity{
		Codepoints: []int64{1013},
		Characters: "\u03f5",
	},
	"&emptysmallsquare;": Entity{
		Codepoints: []int64{9723},
		Characters: "\u25fb",
	},
	"&rtimes;": Entity{
		Codepoints: []int64{8906},
		Characters: "\u22ca",
	},
	"&nang;": Entity{
		Codepoints: []int64{8736, 8402},
		Characters: "\u2220\u20d2",
	},
	"&capcap;": Entity{
		Codepoints: []int64{10827},
		Characters: "\u2a4b",
	},
	"&solb;": Entity{
		Codepoints: []int64{10692},
		Characters: "\u29c4",
	},
	"&ap;": Entity{
		Codepoints: []int64{8776},
		Characters: "\u2248",
	},
	"&exist;": Entity{
		Codepoints: []int64{8707},
		Characters: "\u2203",
	},
	"&rhar;": Entity{
		Codepoints: []int64{10596},
		Characters: "\u2964",
	},
	"&frac38;": Entity{
		Codepoints: []int64{8540},
		Characters: "\u215c",
	},
	"&hcirc;": Entity{
		Codepoints: []int64{293},
		Characters: "\u0125",
	},
	"&ufr;": Entity{
		Codepoints: []int64{120114},
		Characters: "\U0001d532",
	},
	"&frasl;": Entity{
		Codepoints: []int64{8260},
		Characters: "\u2044",
	},
	"&notrighttrianglebar;": Entity{
		Codepoints: []int64{10704, 824},
		Characters: "\u29d0\u0338",
	},
	"&minusd;": Entity{
		Codepoints: []int64{8760},
		Characters: "\u2238",
	},
	"&leftupvectorbar;": Entity{
		Codepoints: []int64{10584},
		Characters: "\u2958",
	},
	"&infin;": Entity{
		Codepoints: []int64{8734},
		Characters: "\u221e",
	},
	"&mid;": Entity{
		Codepoints: []int64{8739},
		Characters: "\u2223",
	},
	"&angrtvbd;": Entity{
		Codepoints: []int64{10653},
		Characters: "\u299d",
	},
	"&ldquo;": Entity{
		Codepoints: []int64{8220},
		Characters: "\u201c",
	},
	"&wedgeq;": Entity{
		Codepoints: []int64{8793},
		Characters: "\u2259",
	},
	"&uuml;": Entity{
		Codepoints: []int64{220},
		Characters: "\u00dc",
	},
	"&rnmid;": Entity{
		Codepoints: []int64{10990},
		Characters: "\u2aee",
	},
	"&coprod;": Entity{
		Codepoints: []int64{8720},
		Characters: "\u2210",
	},
	"&numero;": Entity{
		Codepoints: []int64{8470},
		Characters: "\u2116",
	},
	"&supmult;": Entity{
		Codepoints: []int64{10946},
		Characters: "\u2ac2",
	},
	"&kopf;": Entity{
		Codepoints: []int64{120130},
		Characters: "\U0001d542",
	},
	"&ominus;": Entity{
		Codepoints: []int64{8854},
		Characters: "\u2296",
	},
	"&frac35;": Entity{
		Codepoints: []int64{8535},
		Characters: "\u2157",
	},
	"&ntgl;": Entity{
		Codepoints: []int64{8825},
		Characters: "\u2279",
	},
	"&ulcrop;": Entity{
		Codepoints: []int64{8975},
		Characters: "\u230f",
	},
	"&jscr;": Entity{
		Codepoints: []int64{119973},
		Characters: "\U0001d4a5",
	},
	"&darr;": Entity{
		Codepoints: []int64{8659},
		Characters: "\u21d3",
	},
	"&ffr;": Entity{
		Codepoints: []int64{120099},
		Characters: "\U0001d523",
	},
	"&ntilde;": Entity{
		Codepoints: []int64{241},
		Characters: "\u00f1",
	},
	"&sup1;": Entity{
		Codepoints: []int64{185},
		Characters: "\u00b9",
	},
	"&cupdot;": Entity{
		Codepoints: []int64{8845},
		Characters: "\u228d",
	},
	"&nwnear;": Entity{
		Codepoints: []int64{10535},
		Characters: "\u2927",
	},
	"&fcy;": Entity{
		Codepoints: []int64{1092},
		Characters: "\u0444",
	},
	"&epar;": Entity{
		Codepoints: []int64{8917},
		Characters: "\u22d5",
	},
	"&male;": Entity{
		Codepoints: []int64{9794},
		Characters: "\u2642",
	},
	"&acirc;": Entity{
		Codepoints: []int64{226},
		Characters: "\u00e2",
	},
	"&piv;": Entity{
		Codepoints: []int64{982},
		Characters: "\u03d6",
	},
	"&cuepr;": Entity{
		Codepoints: []int64{8926},
		Characters: "\u22de",
	},
	"&vcy;": Entity{
		Codepoints: []int64{1074},
		Characters: "\u0432",
	},
	"&isins;": Entity{
		Codepoints: []int64{8948},
		Characters: "\u22f4",
	},
	"&roang;": Entity{
		Codepoints: []int64{10221},
		Characters: "\u27ed",
	},
	"&xscr;": Entity{
		Codepoints: []int64{119987},
		Characters: "\U0001d4b3",
	},
	"&rdquo;": Entity{
		Codepoints: []int64{8221},
		Characters: "\u201d",
	},
	"&tridot;": Entity{
		Codepoints: []int64{9708},
		Characters: "\u25ec",
	},
	"&equest;": Entity{
		Codepoints: []int64{8799},
		Characters: "\u225f",
	},
	"&pre;": Entity{
		Codepoints: []int64{10927},
		Characters: "\u2aaf",
	},
	"&curarrm;": Entity{
		Codepoints: []int64{10556},
		Characters: "\u293c",
	},
	"&leftvectorbar;": Entity{
		Codepoints: []int64{10578},
		Characters: "\u2952",
	},
	"&lscr;": Entity{
		Codepoints: []int64{120001},
		Characters: "\U0001d4c1",
	},
	"&pertenk;": Entity{
		Codepoints: []int64{8241},
		Characters: "\u2031",
	},
	"&nprcue;": Entity{
		Codepoints: []int64{8928},
		Characters: "\u22e0",
	},
	"&iuml;": Entity{
		Codepoints: []int64{239},
		Characters: "\u00ef",
	},
	"&sharp;": Entity{
		Codepoints: []int64{9839},
		Characters: "\u266f",
	},
	"&capbrcup;": Entity{
		Codepoints: []int64{10825},
		Characters: "\u2a49",
	},
	"&subsub;": Entity{
		Codepoints: []int64{10965},
		Characters: "\u2ad5",
	},
	"&duarr;": Entity{
		Codepoints: []int64{8693},
		Characters: "\u21f5",
	},
	"&supdot;": Entity{
		Codepoints: []int64{10942},
		Characters: "\u2abe",
	},
	"&yacute;": Entity{
		Codepoints: []int64{253},
		Characters: "\u00fd",
	},
	"&disin;": Entity{
		Codepoints: []int64{8946},
		Characters: "\u22f2",
	},
	"&iexcl;": Entity{
		Codepoints: []int64{161},
		Characters: "\u00a1",
	},
	"&zacute;": Entity{
		Codepoints: []int64{378},
		Characters: "\u017a",
	},
	"&lhar;": Entity{
		Codepoints: []int64{10594},
		Characters: "\u2962",
	},
	"&permil;": Entity{
		Codepoints: []int64{8240},
		Characters: "\u2030",
	},
	"&prap;": Entity{
		Codepoints: []int64{10935},
		Characters: "\u2ab7",
	},
	"&rdsh;": Entity{
		Codepoints: []int64{8627},
		Characters: "\u21b3",
	},
	"&yacy;": Entity{
		Codepoints: []int64{1103},
		Characters: "\u044f",
	},
	"&shcy;": Entity{
		Codepoints: []int64{1064},
		Characters: "\u0428",
	},
	"&ropar;": Entity{
		Codepoints: []int64{10630},
		Characters: "\u2986",
	},
	"&xfr;": Entity{
		Codepoints: []int64{120091},
		Characters: "\U0001d51b",
	},
	"&rsquo;": Entity{
		Codepoints: []int64{8217},
		Characters: "\u2019",
	},
	"&lmidot;": Entity{
		Codepoints: []int64{319},
		Characters: "\u013f",
	},
	"&cemptyv;": Entity{
		Codepoints: []int64{10674},
		Characters: "\u29b2",
	},
	"&lap;": Entity{
		Codepoints: []int64{10885},
		Characters: "\u2a85",
	},
	"&scap;": Entity{
		Codepoints: []int64{10936},
		Characters: "\u2ab8",
	},
	"&lvne;": Entity{
		Codepoints: []int64{8808, 65024},
		Characters: "\u2268\ufe00",
	},
	"&ograve;": Entity{
		Codepoints: []int64{210},
		Characters: "\u00d2",
	},
	"&larrpl;": Entity{
		Codepoints: []int64{10553},
		Characters: "\u2939",
	},
	"&nrtri;": Entity{
		Codepoints: []int64{8939},
		Characters: "\u22eb",
	},
	"&zerowidthspace;": Entity{
		Codepoints: []int64{8203},
		Characters: "\u200b",
	},
	"&pm;": Entity{
		Codepoints: []int64{177},
		Characters: "\u00b1",
	},
	"&gtrarr;": Entity{
		Codepoints: []int64{10616},
		Characters: "\u2978",
	},
	"&yscr;": Entity{
		Codepoints: []int64{119988},
		Characters: "\U0001d4b4",
	},
	"&nbump;": Entity{
		Codepoints: []int64{8782, 824},
		Characters: "\u224e\u0338",
	},
	"&mopf;": Entity{
		Codepoints: []int64{120158},
		Characters: "\U0001d55e",
	},
	"&emsp;": Entity{
		Codepoints: []int64{8195},
		Characters: "\u2003",
	},
	"&sdotb;": Entity{
		Codepoints: []int64{8865},
		Characters: "\u22a1",
	},
	"&angrt;": Entity{
		Codepoints: []int64{8735},
		Characters: "\u221f",
	},
	"&ltrie;": Entity{
		Codepoints: []int64{8884},
		Characters: "\u22b4",
	},
	"&xuplus;": Entity{
		Codepoints: []int64{10756},
		Characters: "\u2a04",
	},
	"&gcirc;": Entity{
		Codepoints: []int64{285},
		Characters: "\u011d",
	},
	"&rrarr;": Entity{
		Codepoints: []int64{8649},
		Characters: "\u21c9",
	},
	"&downarrowbar;": Entity{
		Codepoints: []int64{10515},
		Characters: "\u2913",
	},
	"&uhblk;": Entity{
		Codepoints: []int64{9600},
		Characters: "\u2580",
	},
	"&boxur;": Entity{
		Codepoints: []int64{9492},
		Characters: "\u2514",
	},
	"&dtdot;": Entity{
		Codepoints: []int64{8945},
		Characters: "\u22f1",
	},
	"&scnap;": Entity{
		Codepoints: []int64{10938},
		Characters: "\u2aba",
	},
	"&rharu;": Entity{
		Codepoints: []int64{8640},
		Characters: "\u21c0",
	},
	"&urcorn;": Entity{
		Codepoints: []int64{8989},
		Characters: "\u231d",
	},
	"&cconint;": Entity{
		Codepoints: []int64{8752},
		Characters: "\u2230",
	},
	"&uarrocir;": Entity{
		Codepoints: []int64{10569},
		Characters: "\u2949",
	},
	"&clubs;": Entity{
		Codepoints: []int64{9827},
		Characters: "\u2663",
	},
	"&oline;": Entity{
		Codepoints: []int64{8254},
		Characters: "\u203e",
	},
	"&nlsim;": Entity{
		Codepoints: []int64{8820},
		Characters: "\u2274",
	},
	"&rsqb;": Entity{
		Codepoints: []int64{93},
		Characters: "]",
	},
	"&half;": Entity{
		Codepoints: []int64{189},
		Characters: "\u00bd",
	},
	"&sigma;": Entity{
		Codepoints: []int64{963},
		Characters: "\u03c3",
	},
	"&rlhar;": Entity{
		Codepoints: []int64{8652},
		Characters: "\u21cc",
	},
	"&zcaron;": Entity{
		Codepoints: []int64{382},
		Characters: "\u017e",
	},
	"&edot;": Entity{
		Codepoints: []int64{279},
		Characters: "\u0117",
	},
	"&horbar;": Entity{
		Codepoints: []int64{8213},
		Characters: "\u2015",
	},
	"&mlcp;": Entity{
		Codepoints: []int64{10971},
		Characters: "\u2adb",
	},
	"&gesdotol;": Entity{
		Codepoints: []int64{10884},
		Characters: "\u2a84",
	},
	"&scsim;": Entity{
		Codepoints: []int64{8831},
		Characters: "\u227f",
	},
	"&nsupe;": Entity{
		Codepoints: []int64{8841},
		Characters: "\u2289",
	},
	"&triminus;": Entity{
		Codepoints: []int64{10810},
		Characters: "\u2a3a",
	},
	"&gtquest;": Entity{
		Codepoints: []int64{10876},
		Characters: "\u2a7c",
	},
	"&nvge;": Entity{
		Codepoints: []int64{8805, 8402},
		Characters: "\u2265\u20d2",
	},
	"&xi;": Entity{
		Codepoints: []int64{926},
		Characters: "\u039e",
	},
	"&rthree;": Entity{
		Codepoints: []int64{8908},
		Characters: "\u22cc",
	},
	"&ltcir;": Entity{
		Codepoints: []int64{10873},
		Characters: "\u2a79",
	},
	"&eparsl;": Entity{
		Codepoints: []int64{10723},
		Characters: "\u29e3",
	},
	"&iota;": Entity{
		Codepoints: []int64{953},
		Characters: "\u03b9",
	},
	"&verticalseparator;": Entity{
		Codepoints: []int64{10072},
		Characters: "\u2758",
	},
	"&loarr;": Entity{
		Codepoints: []int64{8701},
		Characters: "\u21fd",
	},
	"&ncedil;": Entity{
		Codepoints: []int64{325},
		Characters: "\u0145",
	},
	"&oslash;": Entity{
		Codepoints: []int64{216},
		Characters: "\u00d8",
	},
	"&equal;": Entity{
		Codepoints: []int64{10869},
		Characters: "\u2a75",
	},
	"&tint;": Entity{
		Codepoints: []int64{8749},
		Characters: "\u222d",
	},
	"&female;": Entity{
		Codepoints: []int64{9792},
		Characters: "\u2640",
	},
	"&csup;": Entity{
		Codepoints: []int64{10960},
		Characters: "\u2ad0",
	},
	"&notni;": Entity{
		Codepoints: []int64{8716},
		Characters: "\u220c",
	},
	"&homtht;": Entity{
		Codepoints: []int64{8763},
		Characters: "\u223b",
	},
	"&nis;": Entity{
		Codepoints: []int64{8956},
		Characters: "\u22fc",
	},
	"&supsim;": Entity{
		Codepoints: []int64{10952},
		Characters: "\u2ac8",
	},
	"&loz;": Entity{
		Codepoints: []int64{9674},
		Characters: "\u25ca",
	},
	"&ges;": Entity{
		Codepoints: []int64{10878},
		Characters: "\u2a7e",
	},
	"&cupor;": Entity{
		Codepoints: []int64{10821},
		Characters: "\u2a45",
	},
	"&iquest;": Entity{
		Codepoints: []int64{191},
		Characters: "\u00bf",
	},
	"&raarr;": Entity{
		Codepoints: []int64{8667},
		Characters: "\u21db",
	},
	"&fllig;": Entity{
		Codepoints: []int64{64258},
		Characters: "\ufb02",
	},
	"&diams;": Entity{
		Codepoints: []int64{9830},
		Characters: "\u2666",
	},
	"&rfisht;": Entity{
		Codepoints: []int64{10621},
		Characters: "\u297d",
	},
	"&fork;": Entity{
		Codepoints: []int64{8916},
		Characters: "\u22d4",
	},
	"&gacute;": Entity{
		Codepoints: []int64{501},
		Characters: "\u01f5",
	},
	"&middot;": Entity{
		Codepoints: []int64{183},
		Characters: "\u00b7",
	},
	"&circ;": Entity{
		Codepoints: []int64{710},
		Characters: "\u02c6",
	},
	"&ultri;": Entity{
		Codepoints: []int64{9720},
		Characters: "\u25f8",
	},
	"&nvlarr;": Entity{
		Codepoints: []int64{10498},
		Characters: "\u2902",
	},
	"&iecy;": Entity{
		Codepoints: []int64{1077},
		Characters: "\u0435",
	},
	"&cularrp;": Entity{
		Codepoints: []int64{10557},
		Characters: "\u293d",
	},
	"&nsup;": Entity{
		Codepoints: []int64{8837},
		Characters: "\u2285",
	},
	"&oacute;": Entity{
		Codepoints: []int64{211},
		Characters: "\u00d3",
	},
	"&notsquaresuperset;": Entity{
		Codepoints: []int64{8848, 824},
		Characters: "\u2290\u0338",
	},
	"&qscr;": Entity{
		Codepoints: []int64{120006},
		Characters: "\U0001d4c6",
	},
	"&mscr;": Entity{
		Codepoints: []int64{8499},
		Characters: "\u2133",
	},
	"&rcaron;": Entity{
		Codepoints: []int64{344},
		Characters: "\u0158",
	},
	"&ic;": Entity{
		Codepoints: []int64{8291},
		Characters: "\u2063",
	},
	"&frac34;": Entity{
		Codepoints: []int64{190},
		Characters: "\u00be",
	},
	"&frac56;": Entity{
		Codepoints: []int64{8538},
		Characters: "\u215a",
	},
	"&ordf;": Entity{
		Codepoints: []int64{170},
		Characters: "\u00aa",
	},
	"&profalar;": Entity{
		Codepoints: []int64{9006},
		Characters: "\u232e",
	},
	"&leftteevector;": Entity{
		Codepoints: []int64{10586},
		Characters: "\u295a",
	},
	"&rightdownteevector;": Entity{
		Codepoints: []int64{10589},
		Characters: "\u295d",
	},
	"&ange;": Entity{
		Codepoints: []int64{10660},
		Characters: "\u29a4",
	},
	"&gsim;": Entity{
		Codepoints: []int64{8819},
		Characters: "\u2273",
	},
	"&capdot;": Entity{
		Codepoints: []int64{10816},
		Characters: "\u2a40",
	},
	"&apos;": Entity{
		Codepoints: []int64{39},
		Characters: "'",
	},
	"&cupcup;": Entity{
		Codepoints: []int64{10826},
		Characters: "\u2a4a",
	},
	"&kcedil;": Entity{
		Codepoints: []int64{310},
		Characters: "\u0136",
	},
	"&lbrksld;": Entity{
		Codepoints: []int64{10639},
		Characters: "\u298f",
	},
	"&vnsup;": Entity{
		Codepoints: []int64{8835, 8402},
		Characters: "\u2283\u20d2",
	},
	"&models;": Entity{
		Codepoints: []int64{8871},
		Characters: "\u22a7",
	},
	"&uwangle;": Entity{
		Codepoints: []int64{10663},
		Characters: "\u29a7",
	},
	"&newline;": Entity{
		Codepoints: []int64{10},
		Characters: "\n",
	},
	"&squ;": Entity{
		Codepoints: []int64{9633},
		Characters: "\u25a1",
	},
	"&ang;": Entity{
		Codepoints: []int64{8736},
		Characters: "\u2220",
	},
	"&subsim;": Entity{
		Codepoints: []int64{10951},
		Characters: "\u2ac7",
	},
	"&blk14;": Entity{
		Codepoints: []int64{9617},
		Characters: "\u2591",
	},
	"&ropf;": Entity{
		Codepoints: []int64{120163},
		Characters: "\U0001d563",
	},
	"&nles;": Entity{
		Codepoints: []int64{10877, 824},
		Characters: "\u2a7d\u0338",
	},
	"&npar;": Entity{
		Codepoints: []int64{8742},
		Characters: "\u2226",
	},
	"&sol;": Entity{
		Codepoints: []int64{47},
		Characters: "/",
	},
	"&lthree;": Entity{
		Codepoints: []int64{8907},
		Characters: "\u22cb",
	},
	"&odblac;": Entity{
		Codepoints: []int64{337},
		Characters: "\u0151",
	},
	"&ape;": Entity{
		Codepoints: []int64{8778},
		Characters: "\u224a",
	},
	"&rightupvectorbar;": Entity{
		Codepoints: []int64{10580},
		Characters: "\u2954",
	},
	"&pfr;": Entity{
		Codepoints: []int64{120109},
		Characters: "\U0001d52d",
	},
	"&bsol;": Entity{
		Codepoints: []int64{92},
		Characters: "\\",
	},
	"&ngt;": Entity{
		Codepoints: []int64{8815},
		Characters: "\u226f",
	},
	"&hardcy;": Entity{
		Codepoints: []int64{1066},
		Characters: "\u042a",
	},
	"&spades;": Entity{
		Codepoints: []int64{9824},
		Characters: "\u2660",
	},
	"&mddot;": Entity{
		Codepoints: []int64{8762},
		Characters: "\u223a",
	},
	"&hat;": Entity{
		Codepoints: []int64{94},
		Characters: "^",
	},
	"&ncaron;": Entity{
		Codepoints: []int64{327},
		Characters: "\u0147",
	},
	"&roundimplies;": Entity{
		Codepoints: []int64{10608},
		Characters: "\u2970",
	},
	"&angmsdab;": Entity{
		Codepoints: []int64{10665},
		Characters: "\u29a9",
	},
	"&incare;": Entity{
		Codepoints: []int64{8453},
		Characters: "\u2105",
	},
	"&cupcap;": Entity{
		Codepoints: []int64{10822},
		Characters: "\u2a46",
	},
	"&luruhar;": Entity{
		Codepoints: []int64{10598},
		Characters: "\u2966",
	},
	"&gdot;": Entity{
		Codepoints: []int64{289},
		Characters: "\u0121",
	},
	"&rarrhk;": Entity{
		Codepoints: []int64{8618},
		Characters: "\u21aa",
	},
	"&vsupne;": Entity{
		Codepoints: []int64{10956, 65024},
		Characters: "\u2acc\ufe00",
	},
	"&rbarr;": Entity{
		Codepoints: []int64{10509},
		Characters: "\u290d",
	},
	"&isindot;": Entity{
		Codepoints: []int64{8949},
		Characters: "\u22f5",
	},
	"&boxvh;": Entity{
		Codepoints: []int64{9578},
		Characters: "\u256a",
	},
	"&supe;": Entity{
		Codepoints: []int64{8839},
		Characters: "\u2287",
	},
	"&theta;": Entity{
		Codepoints: []int64{920},
		Characters: "\u0398",
	},
	"&rarrpl;": Entity{
		Codepoints: []int64{10565},
		Characters: "\u2945",
	},
	"&ouml;": Entity{
		Codepoints: []int64{246},
		Characters: "\u00f6",
	},
	"&ltlarr;": Entity{
		Codepoints: []int64{10614},
		Characters: "\u2976",
	},
	"&caps;": Entity{
		Codepoints: []int64{8745, 65024},
		Characters: "\u2229\ufe00",
	},
	"&sqcap;": Entity{
		Codepoints: []int64{8851},
		Characters: "\u2293",
	},
	"&nvle;": Entity{
		Codepoints: []int64{8804, 8402},
		Characters: "\u2264\u20d2",
	},
	"&larrsim;": Entity{
		Codepoints: []int64{10611},
		Characters: "\u2973",
	},
	"&gtcc;": Entity{
		Codepoints: []int64{10919},
		Characters: "\u2aa7",
	},
	"&lrtri;": Entity{
		Codepoints: []int64{8895},
		Characters: "\u22bf",
	},
	"&csub;": Entity{
		Codepoints: []int64{10959},
		Characters: "\u2acf",
	},
	"&rightteevector;": Entity{
		Codepoints: []int64{10587},
		Characters: "\u295b",
	},
	"&compfn;": Entity{
		Codepoints: []int64{8728},
		Characters: "\u2218",
	},
	"&dcy;": Entity{
		Codepoints: []int64{1076},
		Characters: "\u0434",
	},
	"&glj;": Entity{
		Codepoints: []int64{10916},
		Characters: "\u2aa4",
	},
	"&nvap;": Entity{
		Codepoints: []int64{8781, 8402},
		Characters: "\u224d\u20d2",
	},
	"&submult;": Entity{
		Codepoints: []int64{10945},
		Characters: "\u2ac1",
	},
	"&oror;": Entity{
		Codepoints: []int64{10838},
		Characters: "\u2a56",
	},
	"&cwconint;": Entity{
		Codepoints: []int64{8754},
		Characters: "\u2232",
	},
	"&dstrok;": Entity{
		Codepoints: []int64{272},
		Characters: "\u0110",
	},
	"&iogon;": Entity{
		Codepoints: []int64{303},
		Characters: "\u012f",
	},
	"&lopf;": Entity{
		Codepoints: []int64{120131},
		Characters: "\U0001d543",
	},
	"&sqrt;": Entity{
		Codepoints: []int64{8730},
		Characters: "\u221a",
	},
	"&lmoust;": Entity{
		Codepoints: []int64{9136},
		Characters: "\u23b0",
	},
	"&otimes;": Entity{
		Codepoints: []int64{10807},
		Characters: "\u2a37",
	},
	"&lang;": Entity{
		Codepoints: []int64{10216},
		Characters: "\u27e8",
	},
	"&jsercy;": Entity{
		Codepoints: []int64{1112},
		Characters: "\u0458",
	},
	"&bdquo;": Entity{
		Codepoints: []int64{8222},
		Characters: "\u201e",
	},
	"&imped;": Entity{
		Codepoints: []int64{437},
		Characters: "\u01b5",
	},
	"&frac45;": Entity{
		Codepoints: []int64{8536},
		Characters: "\u2158",
	},
	"&andv;": Entity{
		Codepoints: []int64{10842},
		Characters: "\u2a5a",
	},
	"&plusacir;": Entity{
		Codepoints: []int64{10787},
		Characters: "\u2a23",
	},
	"&oopf;": Entity{
		Codepoints: []int64{120160},
		Characters: "\U0001d560",
	},
	"&nltv;": Entity{
		Codepoints: []int64{8810, 824},
		Characters: "\u226a\u0338",
	},
	"&trisb;": Entity{
		Codepoints: []int64{10701},
		Characters: "\u29cd",
	},
	"&sbquo;": Entity{
		Codepoints: []int64{8218},
		Characters: "\u201a",
	},
	"&cirfnint;": Entity{
		Codepoints: []int64{10768},
		Characters: "\u2a10",
	},
	"&subedot;": Entity{
		Codepoints: []int64{10947},
		Characters: "\u2ac3",
	},
	"&bfr;": Entity{
		Codepoints: []int64{120069},
		Characters: "\U0001d505",
	},
	"&mumap;": Entity{
		Codepoints: []int64{8888},
		Characters: "\u22b8",
	},
	"&minus;": Entity{
		Codepoints: []int64{8722},
		Characters: "\u2212",
	},
	"&chcy;": Entity{
		Codepoints: []int64{1095},
		Characters: "\u0447",
	},
	"&uopf;": Entity{
		Codepoints: []int64{120166},
		Characters: "\U0001d566",
	},
	"&boxh;": Entity{
		Codepoints: []int64{9472},
		Characters: "\u2500",
	},
	"&bprime;": Entity{
		Codepoints: []int64{8245},
		Characters: "\u2035",
	},
	"&cups;": Entity{
		Codepoints: []int64{8746, 65024},
		Characters: "\u222a\ufe00",
	},
	"&simge;": Entity{
		Codepoints: []int64{10912},
		Characters: "\u2aa0",
	},
	"&laarr;": Entity{
		Codepoints: []int64{8666},
		Characters: "\u21da",
	},
	"&lhard;": Entity{
		Codepoints: []int64{8637},
		Characters: "\u21bd",
	},
	"&yuml;": Entity{
		Codepoints: []int64{376},
		Characters: "\u0178",
	},
	"&boxv;": Entity{
		Codepoints: []int64{9553},
		Characters: "\u2551",
	},
	"&ruluhar;": Entity{
		Codepoints: []int64{10600},
		Characters: "\u2968",
	},
	"&gl;": Entity{
		Codepoints: []int64{8823},
		Characters: "\u2277",
	},
	"&ltrpar;": Entity{
		Codepoints: []int64{10646},
		Characters: "\u2996",
	},
	"&nvharr;": Entity{
		Codepoints: []int64{10500},
		Characters: "\u2904",
	},
	"&npre;": Entity{
		Codepoints: []int64{10927, 824},
		Characters: "\u2aaf\u0338",
	},
	"&bne;": Entity{
		Codepoints: []int64{61, 8421},
		Characters: "=\u20e5",
	},
	"&vscr;": Entity{
		Codepoints: []int64{120011},
		Characters: "\U0001d4cb",
	},
	"&xsqcup;": Entity{
		Codepoints: []int64{10758},
		Characters: "\u2a06",
	},
	"&plus;": Entity{
		Codepoints: []int64{43},
		Characters: "+",
	},
	"&omacr;": Entity{
		Codepoints: []int64{332},
		Characters: "\u014c",
	},
	"&sup;": Entity{
		Codepoints: []int64{8835},
		Characters: "\u2283",
	},
	"&parsim;": Entity{
		Codepoints: []int64{10995},
		Characters: "\u2af3",
	},
	"&lotimes;": Entity{
		Codepoints: []int64{10804},
		Characters: "\u2a34",
	},
	"&nbsp;": Entity{
		Codepoints: []int64{160},
		Characters: "\u00a0",
	},
	"&urtri;": Entity{
		Codepoints: []int64{9721},
		Characters: "\u25f9",
	},
	"&ucirc;": Entity{
		Codepoints: []int64{219},
		Characters: "\u00db",
	},
	"&bnequiv;": Entity{
		Codepoints: []int64{8801, 8421},
		Characters: "\u2261\u20e5",
	},
	"&lesges;": Entity{
		Codepoints: []int64{10899},
		Characters: "\u2a93",
	},
	"&curren;": Entity{
		Codepoints: []int64{164},
		Characters: "\u00a4",
	},
	"&rightupteevector;": Entity{
		Codepoints: []int64{10588},
		Characters: "\u295c",
	},
	"&emsp13;": Entity{
		Codepoints: []int64{8196},
		Characters: "\u2004",
	},
	"&notsucceedstilde;": Entity{
		Codepoints: []int64{8831, 824},
		Characters: "\u227f\u0338",
	},
	"&zhcy;": Entity{
		Codepoints: []int64{1078},
		Characters: "\u0436",
	},
	"&lrarr;": Entity{
		Codepoints: []int64{8646},
		Characters: "\u21c6",
	},
	"&shchcy;": Entity{
		Codepoints: []int64{1065},
		Characters: "\u0429",
	},
	"&odot;": Entity{
		Codepoints: []int64{8857},
		Characters: "\u2299",
	},
	"&angrtvb;": Entity{
		Codepoints: []int64{8894},
		Characters: "\u22be",
	},
	"&eogon;": Entity{
		Codepoints: []int64{280},
		Characters: "\u0118",
	},
	"&scaron;": Entity{
		Codepoints: []int64{353},
		Characters: "\u0161",
	},
	"&raquo;": Entity{
		Codepoints: []int64{187},
		Characters: "\u00bb",
	},
	"&ccupssm;": Entity{
		Codepoints: []int64{10832},
		Characters: "\u2a50",
	},
	"&prnsim;": Entity{
		Codepoints: []int64{8936},
		Characters: "\u22e8",
	},
	"&rpargt;": Entity{
		Codepoints: []int64{10644},
		Characters: "\u2994",
	},
	"&mapstoup;": Entity{
		Codepoints: []int64{8613},
		Characters: "\u21a5",
	},
	"&trpezium;": Entity{
		Codepoints: []int64{9186},
		Characters: "\u23e2",
	},
	"&ngg;": Entity{
		Codepoints: []int64{8921, 824},
		Characters: "\u22d9\u0338",
	},
	"&crarr;": Entity{
		Codepoints: []int64{8629},
		Characters: "\u21b5",
	},
	"&larrbfs;": Entity{
		Codepoints: []int64{10527},
		Characters: "\u291f",
	},
	"&mediumspace;": Entity{
		Codepoints: []int64{8287},
		Characters: "\u205f",
	},
	"&rhov;": Entity{
		Codepoints: []int64{1009},
		Characters: "\u03f1",
	},
	"&oast;": Entity{
		Codepoints: []int64{8859},
		Characters: "\u229b",
	},
	"&notcupcap;": Entity{
		Codepoints: []int64{8813},
		Characters: "\u226d",
	},
	"&rarrlp;": Entity{
		Codepoints: []int64{8620},
		Characters: "\u21ac",
	},
	"&oint;": Entity{
		Codepoints: []int64{8750},
		Characters: "\u222e",
	},
	"&eacute;": Entity{
		Codepoints: []int64{201},
		Characters: "\u00c9",
	},
	"&duhar;": Entity{
		Codepoints: []int64{10607},
		Characters: "\u296f",
	},
	"&ngsim;": Entity{
		Codepoints: []int64{8821},
		Characters: "\u2275",
	},
	"&quot;": Entity{
		Codepoints: []int64{34},
		Characters: "\"",
	},
	"&csupe;": Entity{
		Codepoints: []int64{10962},
		Characters: "\u2ad2",
	},
	"&rarrb;": Entity{
		Codepoints: []int64{8677},
		Characters: "\u21e5",
	},
	"&hacek;": Entity{
		Codepoints: []int64{711},
		Characters: "\u02c7",
	},
	"&notine;": Entity{
		Codepoints: []int64{8953, 824},
		Characters: "\u22f9\u0338",
	},
	"&copy;": Entity{
		Codepoints: []int64{169},
		Characters: "\u00a9",
	},
	"&iukcy;": Entity{
		Codepoints: []int64{1110},
		Characters: "\u0456",
	},
	"&blk12;": Entity{
		Codepoints: []int64{9618},
		Characters: "\u2592",
	},
	"&doteq;": Entity{
		Codepoints: []int64{8784},
		Characters: "\u2250",
	},
	"&topf;": Entity{
		Codepoints: []int64{120139},
		Characters: "\U0001d54b",
	},
	"&hercon;": Entity{
		Codepoints: []int64{8889},
		Characters: "\u22b9",
	},
	"&vrtri;": Entity{
		Codepoints: []int64{8883},
		Characters: "\u22b3",
	},
	"&angzarr;": Entity{
		Codepoints: []int64{9084},
		Characters: "\u237c",
	},
	"&eta;": Entity{
		Codepoints: []int64{919},
		Characters: "\u0397",
	},
	"&bsemi;": Entity{
		Codepoints: []int64{8271},
		Characters: "\u204f",
	},
	"&ecy;": Entity{
		Codepoints: []int64{1101},
		Characters: "\u044d",
	},
	"&wopf;": Entity{
		Codepoints: []int64{120168},
		Characters: "\U0001d568",
	},
	"&scedil;": Entity{
		Codepoints: []int64{350},
		Characters: "\u015e",
	},
	"&scy;": Entity{
		Codepoints: []int64{1089},
		Characters: "\u0441",
	},
	"&rbbrk;": Entity{
		Codepoints: []int64{10099},
		Characters: "\u2773",
	},
	"&gjcy;": Entity{
		Codepoints: []int64{1107},
		Characters: "\u0453",
	},
	"&bcy;": Entity{
		Codepoints: []int64{1073},
		Characters: "\u0431",
	},
	"&ofr;": Entity{
		Codepoints: []int64{120108},
		Characters: "\U0001d52c",
	},
	"&varr;": Entity{
		Codepoints: []int64{8661},
		Characters: "\u21d5",
	},
	"&cwint;": Entity{
		Codepoints: []int64{8753},
		Characters: "\u2231",
	},
	"&empty;": Entity{
		Codepoints: []int64{8709},
		Characters: "\u2205",
	},
	"&qfr;": Entity{
		Codepoints: []int64{120110},
		Characters: "\U0001d52e",
	},
	"&olcross;": Entity{
		Codepoints: []int64{10683},
		Characters: "\u29bb",
	},
	"&osol;": Entity{
		Codepoints: []int64{8856},
		Characters: "\u2298",
	},
	"&diam;": Entity{
		Codepoints: []int64{8900},
		Characters: "\u22c4",
	},
	"&sect;": Entity{
		Codepoints: []int64{167},
		Characters: "\u00a7",
	},
	"&khcy;": Entity{
		Codepoints: []int64{1061},
		Characters: "\u0425",
	},
	"&phi;": Entity{
		Codepoints: []int64{934},
		Characters: "\u03a6",
	},
	"&mapstodown;": Entity{
		Codepoints: []int64{8615},
		Characters: "\u21a7",
	},
	"&cudarrr;": Entity{
		Codepoints: []int64{10549},
		Characters: "\u2935",
	},
	"&lrhard;": Entity{
		Codepoints: []int64{10605},
		Characters: "\u296d",
	},
	"&equivdd;": Entity{
		Codepoints: []int64{10872},
		Characters: "\u2a78",
	},
	"&bump;": Entity{
		Codepoints: []int64{8782},
		Characters: "\u224e",
	},
	"&eplus;": Entity{
		Codepoints: []int64{10865},
		Characters: "\u2a71",
	},
	"&simdot;": Entity{
		Codepoints: []int64{10858},
		Characters: "\u2a6a",
	},
	"&iiota;": Entity{
		Codepoints: []int64{8489},
		Characters: "\u2129",
	},
	"&loplus;": Entity{
		Codepoints: []int64{10797},
		Characters: "\u2a2d",
	},
	"&sung;": Entity{
		Codepoints: []int64{9834},
		Characters: "\u266a",
	},
	"&egrave;": Entity{
		Codepoints: []int64{200},
		Characters: "\u00c8",
	},
	"&macr;": Entity{
		Codepoints: []int64{175},
		Characters: "\u00af",
	},
	"&ugrave;": Entity{
		Codepoints: []int64{249},
		Characters: "\u00f9",
	},
	"&mu;": Entity{
		Codepoints: []int64{956},
		Characters: "\u03bc",
	},
	"&itilde;": Entity{
		Codepoints: []int64{297},
		Characters: "\u0129",
	},
	"&dblac;": Entity{
		Codepoints: []int64{733},
		Characters: "\u02dd",
	},
	"&notnivc;": Entity{
		Codepoints: []int64{8957},
		Characters: "\u22fd",
	},
	"&leftupdownvector;": Entity{
		Codepoints: []int64{10577},
		Characters: "\u2951",
	},
	"&ifr;": Entity{
		Codepoints: []int64{120102},
		Characters: "\U0001d526",
	},
	"&veeeq;": Entity{
		Codepoints: []int64{8794},
		Characters: "\u225a",
	},
	"&sube;": Entity{
		Codepoints: []int64{8838},
		Characters: "\u2286",
	},
	"&profline;": Entity{
		Codepoints: []int64{8978},
		Characters: "\u2312",
	},
	"&cirmid;": Entity{
		Codepoints: []int64{10991},
		Characters: "\u2aef",
	},
	"&lsquo;": Entity{
		Codepoints: []int64{8216},
		Characters: "\u2018",
	},
	"&fnof;": Entity{
		Codepoints: []int64{402},
		Characters: "\u0192",
	},
	"&cupbrcap;": Entity{
		Codepoints: []int64{10824},
		Characters: "\u2a48",
	},
	"&simplus;": Entity{
		Codepoints: []int64{10788},
		Characters: "\u2a24",
	},
	"&dtrif;": Entity{
		Codepoints: []int64{9662},
		Characters: "\u25be",
	},
	"&wscr;": Entity{
		Codepoints: []int64{119986},
		Characters: "\U0001d4b2",
	},
	"&xutri;": Entity{
		Codepoints: []int64{9651},
		Characters: "\u25b3",
	},
	"&xcup;": Entity{
		Codepoints: []int64{8899},
		Characters: "\u22c3",
	},
	"&lsimg;": Entity{
		Codepoints: []int64{10895},
		Characters: "\u2a8f",
	},
	"&puncsp;": Entity{
		Codepoints: []int64{8200},
		Characters: "\u2008",
	},
	"&dharl;": Entity{
		Codepoints: []int64{8643},
		Characters: "\u21c3",
	},
	"&nsqsupe;": Entity{
		Codepoints: []int64{8931},
		Characters: "\u22e3",
	},
	"&laquo;": Entity{
		Codepoints: []int64{171},
		Characters: "\u00ab",
	},
	"&kfr;": Entity{
		Codepoints: []int64{120104},
		Characters: "\U0001d528",
	},
	"&olarr;": Entity{
		Codepoints: []int64{8634},
		Characters: "\u21ba",
	},
	"&top;": Entity{
		Codepoints: []int64{8868},
		Characters: "\u22a4",
	},
	"&lesdotor;": Entity{
		Codepoints: []int64{10883},
		Characters: "\u2a83",
	},
	"&scpolint;": Entity{
		Codepoints: []int64{10771},
		Characters: "\u2a13",
	},
	"&angmsdae;": Entity{
		Codepoints: []int64{10668},
		Characters: "\u29ac",
	},
	"&lge;": Entity{
		Codepoints: []int64{10897},
		Characters: "\u2a91",
	},
	"&solbar;": Entity{
		Codepoints: []int64{9023},
		Characters: "\u233f",
	},
	"&gamma;": Entity{
		Codepoints: []int64{915},
		Characters: "\u0393",
	},
	"&ncongdot;": Entity{
		Codepoints: []int64{10861, 824},
		Characters: "\u2a6d\u0338",
	},
	"&flat;": Entity{
		Codepoints: []int64{9837},
		Characters: "\u266d",
	},
	"&brvbar;": Entity{
		Codepoints: []int64{166},
		Characters: "\u00a6",
	},
	"&nrarr;": Entity{
		Codepoints: []int64{8655},
		Characters: "\u21cf",
	},
	"&gesles;": Entity{
		Codepoints: []int64{10900},
		Characters: "\u2a94",
	},
	"&xcap;": Entity{
		Codepoints: []int64{8898},
		Characters: "\u22c2",
	},
	"&kjcy;": Entity{
		Codepoints: []int64{1116},
		Characters: "\u045c",
	},
	"&tritime;": Entity{
		Codepoints: []int64{10811},
		Characters: "\u2a3b",
	},
	"&lurdshar;": Entity{
		Codepoints: []int64{10570},
		Characters: "\u294a",
	},
	"&nvinfin;": Entity{
		Codepoints: []int64{10718},
		Characters: "\u29de",
	},
	"&opar;": Entity{
		Codepoints: []int64{10679},
		Characters: "\u29b7",
	},
	"&wedbar;": Entity{
		Codepoints: []int64{10847},
		Characters: "\u2a5f",
	},
	"&omicron;": Entity{
		Codepoints: []int64{927},
		Characters: "\u039f",
	},
	"&larr;": Entity{
		Codepoints: []int64{8606},
		Characters: "\u219e",
	},
	"&nearr;": Entity{
		Codepoints: []int64{8599},
		Characters: "\u2197",
	},
	"&it;": Entity{
		Codepoints: []int64{8290},
		Characters: "\u2062",
	},
	"&xotime;": Entity{
		Codepoints: []int64{10754},
		Characters: "\u2a02",
	},
	"&lcub;": Entity{
		Codepoints: []int64{123},
		Characters: "{",
	},
	"&ntlg;": Entity{
		Codepoints: []int64{8824},
		Characters: "\u2278",
	},
	"&ltcc;": Entity{
		Codepoints: []int64{10918},
		Characters: "\u2aa6",
	},
	"&ratail;": Entity{
		Codepoints: []int64{10522},
		Characters: "\u291a",
	},
	"&ee;": Entity{
		Codepoints: []int64{8519},
		Characters: "\u2147",
	},
	"&ccirc;": Entity{
		Codepoints: []int64{264},
		Characters: "\u0108",
	},
	"&sce;": Entity{
		Codepoints: []int64{10928},
		Characters: "\u2ab0",
	},
	"&lceil;": Entity{
		Codepoints: []int64{8968},
		Characters: "\u2308",
	},
	"&quatint;": Entity{
		Codepoints: []int64{10774},
		Characters: "\u2a16",
	},
	"&qopf;": Entity{
		Codepoints: []int64{8474},
		Characters: "\u211a",
	},
	"&uplus;": Entity{
		Codepoints: []int64{8846},
		Characters: "\u228e",
	},
	"&nltri;": Entity{
		Codepoints: []int64{8938},
		Characters: "\u22ea",
	},
	"&notnestedgreatergreater;": Entity{
		Codepoints: []int64{10914, 824},
		Characters: "\u2aa2\u0338",
	},
	"&bbrk;": Entity{
		Codepoints: []int64{9141},
		Characters: "\u23b5",
	},
	"&ii;": Entity{
		Codepoints: []int64{8520},
		Characters: "\u2148",
	},
	"&tfr;": Entity{
		Codepoints: []int64{120113},
		Characters: "\U0001d531",
	},
	"&beta;": Entity{
		Codepoints: []int64{946},
		Characters: "\u03b2",
	},
	"&uuarr;": Entity{
		Codepoints: []int64{8648},
		Characters: "\u21c8",
	},
	"&overbrace;": Entity{
		Codepoints: []int64{9182},
		Characters: "\u23de",
	},
	"&pluse;": Entity{
		Codepoints: []int64{10866},
		Characters: "\u2a72",
	},
	"&sscr;": Entity{
		Codepoints: []int64{120008},
		Characters: "\U0001d4c8",
	},
	"&rho;": Entity{
		Codepoints: []int64{929},
		Characters: "\u03a1",
	},
	"&lesdot;": Entity{
		Codepoints: []int64{10879},
		Characters: "\u2a7f",
	},
	"&plustwo;": Entity{
		Codepoints: []int64{10791},
		Characters: "\u2a27",
	},
	"&div;": Entity{
		Codepoints: []int64{247},
		Characters: "\u00f7",
	},
	"&ncong;": Entity{
		Codepoints: []int64{8775},
		Characters: "\u2247",
	},
	"&timesbar;": Entity{
		Codepoints: []int64{10801},
		Characters: "\u2a31",
	},
	"&els;": Entity{
		Codepoints: []int64{10901},
		Characters: "\u2a95",
	},
	"&iff;": Entity{
		Codepoints: []int64{8660},
		Characters: "\u21d4",
	},
	"&rarrtl;": Entity{
		Codepoints: []int64{10518},
		Characters: "\u2916",
	},
	"&euro;": Entity{
		Codepoints: []int64{8364},
		Characters: "\u20ac",
	},
	"&re;": Entity{
		Codepoints: []int64{8476},
		Characters: "\u211c",
	},
	"&smte;": Entity{
		Codepoints: []int64{10924},
		Characters: "\u2aac",
	},
	"&le;": Entity{
		Codepoints: []int64{8804},
		Characters: "\u2264",
	},
	"&lnap;": Entity{
		Codepoints: []int64{10889},
		Characters: "\u2a89",
	},
	"&part;": Entity{
		Codepoints: []int64{8706},
		Characters: "\u2202",
	},
	"&blk34;": Entity{
		Codepoints: []int64{9619},
		Characters: "\u2593",
	},
	"&ubreve;": Entity{
		Codepoints: []int64{364},
		Characters: "\u016c",
	},
	"&cudarrl;": Entity{
		Codepoints: []int64{10552},
		Characters: "\u2938",
	},
	"&vltri;": Entity{
		Codepoints: []int64{8882},
		Characters: "\u22b2",
	},
	"&blank;": Entity{
		Codepoints: []int64{9251},
		Characters: "\u2423",
	},
	"&andslope;": Entity{
		Codepoints: []int64{10840},
		Characters: "\u2a58",
	},
	"&angmsdaf;": Entity{
		Codepoints: []int64{10669},
		Characters: "\u29ad",
	},
	"&deg;": Entity{
		Codepoints: []int64{176},
		Characters: "\u00b0",
	},
	"&dagger;": Entity{
		Codepoints: []int64{8225},
		Characters: "\u2021",
	},
	"&udhar;": Entity{
		Codepoints: []int64{10606},
		Characters: "\u296e",
	},
	"&micro;": Entity{
		Codepoints: []int64{181},
		Characters: "\u00b5",
	},
	"&target;": Entity{
		Codepoints: []int64{8982},
		Characters: "\u2316",
	},
	"&esim;": Entity{
		Codepoints: []int64{10867},
		Characters: "\u2a73",
	},
	"&profsurf;": Entity{
		Codepoints: []int64{8979},
		Characters: "\u2313",
	},
	"&gg;": Entity{
		Codepoints: []int64{8811},
		Characters: "\u226b",
	},
	"&easter;": Entity{
		Codepoints: []int64{10862},
		Characters: "\u2a6e",
	},
	"&ndash;": Entity{
		Codepoints: []int64{8211},
		Characters: "\u2013",
	},
	"&prop;": Entity{
		Codepoints: []int64{8733},
		Characters: "\u221d",
	},
	"&simg;": Entity{
		Codepoints: []int64{10910},
		Characters: "\u2a9e",
	},
	"&downrightteevector;": Entity{
		Codepoints: []int64{10591},
		Characters: "\u295f",
	},
	"&boxdr;": Entity{
		Codepoints: []int64{9484},
		Characters: "\u250c",
	},
	"&yopf;": Entity{
		Codepoints: []int64{120144},
		Characters: "\U0001d550",
	},
	"&aopf;": Entity{
		Codepoints: []int64{120146},
		Characters: "\U0001d552",
	},
	"&suphsub;": Entity{
		Codepoints: []int64{10967},
		Characters: "\u2ad7",
	},
	"&ast;": Entity{
		Codepoints: []int64{42},
		Characters: "*",
	},
	"&filig;": Entity{
		Codepoints: []int64{64257},
		Characters: "\ufb01",
	},
	"&mho;": Entity{
		Codepoints: []int64{8487},
		Characters: "\u2127",
	},
	"&tilde;": Entity{
		Codepoints: []int64{732},
		Characters: "\u02dc",
	},
	"&dsol;": Entity{
		Codepoints: []int64{10742},
		Characters: "\u29f6",
	},
	"&lessless;": Entity{
		Codepoints: []int64{10913},
		Characters: "\u2aa1",
	},
	"&tbrk;": Entity{
		Codepoints: []int64{9140},
		Characters: "\u23b4",
	},
	"&gne;": Entity{
		Codepoints: []int64{10888},
		Characters: "\u2a88",
	},
	"&rarrbfs;": Entity{
		Codepoints: []int64{10528},
		Characters: "\u2920",
	},
	"&nparsl;": Entity{
		Codepoints: []int64{11005, 8421},
		Characters: "\u2afd\u20e5",
	},
	"&frac16;": Entity{
		Codepoints: []int64{8537},
		Characters: "\u2159",
	},
	"&iscr;": Entity{
		Codepoints: []int64{8464},
		Characters: "\u2110",
	},
	"&squf;": Entity{
		Codepoints: []int64{9642},
		Characters: "\u25aa",
	},
	"&ne;": Entity{
		Codepoints: []int64{8800},
		Characters: "\u2260",
	},
	"&gesl;": Entity{
		Codepoints: []int64{8923, 65024},
		Characters: "\u22db\ufe00",
	},
	"&nle;": Entity{
		Codepoints: []int64{8806, 824},
		Characters: "\u2266\u0338",
	},
	"&cent;": Entity{
		Codepoints: []int64{162},
		Characters: "\u00a2",
	},
	"&cuesc;": Entity{
		Codepoints: []int64{8927},
		Characters: "\u22df",
	},
	"&dcaron;": Entity{
		Codepoints: []int64{271},
		Characters: "\u010f",
	},
	"&uhar;": Entity{
		Codepoints: []int64{10595},
		Characters: "\u2963",
	},
	"&cirscir;": Entity{
		Codepoints: []int64{10690},
		Characters: "\u29c2",
	},
	"&nsccue;": Entity{
		Codepoints: []int64{8929},
		Characters: "\u22e1",
	},
	"&zfr;": Entity{
		Codepoints: []int64{8488},
		Characters: "\u2128",
	},
	"&percnt;": Entity{
		Codepoints: []int64{37},
		Characters: "%",
	},
	"&lopar;": Entity{
		Codepoints: []int64{10629},
		Characters: "\u2985",
	},
	"&lowast;": Entity{
		Codepoints: []int64{8727},
		Characters: "\u2217",
	},
	"&erdot;": Entity{
		Codepoints: []int64{8787},
		Characters: "\u2253",
	},
	"&divonx;": Entity{
		Codepoints: []int64{8903},
		Characters: "\u22c7",
	},
	"&ocir;": Entity{
		Codepoints: []int64{8858},
		Characters: "\u229a",
	},
	"&dlcrop;": Entity{
		Codepoints: []int64{8973},
		Characters: "\u230d",
	},
	"&boxbox;": Entity{
		Codepoints: []int64{10697},
		Characters: "\u29c9",
	},
	"&xoplus;": Entity{
		Codepoints: []int64{10753},
		Characters: "\u2a01",
	},
	"&angmsdag;": Entity{
		Codepoints: []int64{10670},
		Characters: "\u29ae",
	},
	"&ncup;": Entity{
		Codepoints: []int64{10818},
		Characters: "\u2a42",
	},
	"&curarr;": Entity{
		Codepoints: []int64{8631},
		Characters: "\u21b7",
	},
	"&downbreve;": Entity{
		Codepoints: []int64{785},
		Characters: "\u0311",
	},
	"&rcedil;": Entity{
		Codepoints: []int64{343},
		Characters: "\u0157",
	},
	"&dwangle;": Entity{
		Codepoints: []int64{10662},
		Characters: "\u29a6",
	},
	"&ccaron;": Entity{
		Codepoints: []int64{269},
		Characters: "\u010d",
	},
	"&del;": Entity{
		Codepoints: []int64{8711},
		Characters: "\u2207",
	},
	"&ell;": Entity{
		Codepoints: []int64{8467},
		Characters: "\u2113",
	},
	"&tab;": Entity{
		Codepoints: []int64{9},
		Characters: "\t",
	},
	"&supsup;": Entity{
		Codepoints: []int64{10966},
		Characters: "\u2ad6",
	},
	"&notindot;": Entity{
		Codepoints: []int64{8949, 824},
		Characters: "\u22f5\u0338",
	},
	"&sc;": Entity{
		Codepoints: []int64{8827},
		Characters: "\u227b",
	},
	"&oscr;": Entity{
		Codepoints: []int64{8500},
		Characters: "\u2134",
	},
	"&star;": Entity{
		Codepoints: []int64{9734},
		Characters: "\u2606",
	},
	"&congdot;": Entity{
		Codepoints: []int64{10861},
		Characters: "\u2a6d",
	},
	"&ni;": Entity{
		Codepoints: []int64{8715},
		Characters: "\u220b",
	},
	"&dotdot;": Entity{
		Codepoints: []int64{8412},
		Characters: "\u20dc",
	},
	"&ascr;": Entity{
		Codepoints: []int64{119964},
		Characters: "\U0001d49c",
	},
	"&eg;": Entity{
		Codepoints: []int64{10906},
		Characters: "\u2a9a",
	},
	"&rmoust;": Entity{
		Codepoints: []int64{9137},
		Characters: "\u23b1",
	},
	"&swarr;": Entity{
		Codepoints: []int64{8665},
		Characters: "\u21d9",
	},
	"&equals;": Entity{
		Codepoints: []int64{61},
		Characters: "=",
	},
	"&angsph;": Entity{
		Codepoints: []int64{8738},
		Characters: "\u2222",
	},
	"&downleftrightvector;": Entity{
		Codepoints: []int64{10576},
		Characters: "\u2950",
	},
	"&siml;": Entity{
		Codepoints: []int64{10909},
		Characters: "\u2a9d",
	},
	"&sccue;": Entity{
		Codepoints: []int64{8829},
		Characters: "\u227d",
	},
	"&ljcy;": Entity{
		Codepoints: []int64{1113},
		Characters: "\u0459",
	},
	"&timesd;": Entity{
		Codepoints: []int64{10800},
		Characters: "\u2a30",
	},
	"&thickspace;": Entity{
		Codepoints: []int64{8287, 8202},
		Characters: "\u205f\u200a",
	},
	"&frown;": Entity{
		Codepoints: []int64{8994},
		Characters: "\u2322",
	},
	"&timesb;": Entity{
		Codepoints: []int64{8864},
		Characters: "\u22a0",
	},
	"&vfr;": Entity{
		Codepoints: []int64{120115},
		Characters: "\U0001d533",
	},
	"&sub;": Entity{
		Codepoints: []int64{8912},
		Characters: "\u22d0",
	},
	"&cong;": Entity{
		Codepoints: []int64{8773},
		Characters: "\u2245",
	},
	"&underbrace;": Entity{
		Codepoints: []int64{9183},
		Characters: "\u23df",
	},
	"&mp;": Entity{
		Codepoints: []int64{8723},
		Characters: "\u2213",
	},
	"&xmap;": Entity{
		Codepoints: []int64{10236},
		Characters: "\u27fc",
	},
	"&gla;": Entity{
		Codepoints: []int64{10917},
		Characters: "\u2aa5",
	},
	"&ruledelayed;": Entity{
		Codepoints: []int64{10740},
		Characters: "\u29f4",
	},
	"&latail;": Entity{
		Codepoints: []int64{10523},
		Characters: "\u291b",
	},
	"&dashv;": Entity{
		Codepoints: []int64{10980},
		Characters: "\u2ae4",
	},
	"&lt;": Entity{
		Codepoints: []int64{60},
		Characters: "<",
	},
	"&nedot;": Entity{
		Codepoints: []int64{8784, 824},
		Characters: "\u2250\u0338",
	},
	"&xodot;": Entity{
		Codepoints: []int64{10752},
		Characters: "\u2a00",
	},
	"&les;": Entity{
		Codepoints: []int64{10877},
		Characters: "\u2a7d",
	},
	"&fjlig;": Entity{
		Codepoints: []int64{102, 106},
		Characters: "fj",
	},
	"&sqsupe;": Entity{
		Codepoints: []int64{8850},
		Characters: "\u2292",
	},
	"&lsime;": Entity{
		Codepoints: []int64{10893},
		Characters: "\u2a8d",
	},
	"&pound;": Entity{
		Codepoints: []int64{163},
		Characters: "\u00a3",
	},
	"&iopf;": Entity{
		Codepoints: []int64{120154},
		Characters: "\U0001d55a",
	},
	"&rarrfs;": Entity{
		Codepoints: []int64{10526},
		Characters: "\u291e",
	},
	"&nvdash;": Entity{
		Codepoints: []int64{8877},
		Characters: "\u22ad",
	},
	"&nltrie;": Entity{
		Codepoints: []int64{8940},
		Characters: "\u22ec",
	},
	"&rarrw;": Entity{
		Codepoints: []int64{8605},
		Characters: "\u219d",
	},
	"&dlcorn;": Entity{
		Codepoints: []int64{8990},
		Characters: "\u231e",
	},
	"&roarr;": Entity{
		Codepoints: []int64{8702},
		Characters: "\u21fe",
	},
	"&dtri;": Entity{
		Codepoints: []int64{9663},
		Characters: "\u25bf",
	},
	"&searr;": Entity{
		Codepoints: []int64{8600},
		Characters: "\u2198",
	},
	"&frac25;": Entity{
		Codepoints: []int64{8534},
		Characters: "\u2156",
	},
	"&notinvc;": Entity{
		Codepoints: []int64{8950},
		Characters: "\u22f6",
	},
	"&ordm;": Entity{
		Codepoints: []int64{186},
		Characters: "\u00ba",
	},
	"&ncap;": Entity{
		Codepoints: []int64{10819},
		Characters: "\u2a43",
	},
	"&egs;": Entity{
		Codepoints: []int64{10902},
		Characters: "\u2a96",
	},
	"&rfr;": Entity{
		Codepoints: []int64{120111},
		Characters: "\U0001d52f",
	},
	"&minusb;": Entity{
		Codepoints: []int64{8863},
		Characters: "\u229f",
	},
	"&ufisht;": Entity{
		Codepoints: []int64{10622},
		Characters: "\u297e",
	},
	"&xrarr;": Entity{
		Codepoints: []int64{10230},
		Characters: "\u27f6",
	},
	"&supne;": Entity{
		Codepoints: []int64{10956},
		Characters: "\u2acc",
	},
	"&suphsol;": Entity{
		Codepoints: []int64{10185},
		Characters: "\u27c9",
	},
	"&uacute;": Entity{
		Codepoints: []int64{250},
		Characters: "\u00fa",
	},
	"&supedot;": Entity{
		Codepoints: []int64{10948},
		Characters: "\u2ac4",
	},
	"&capand;": Entity{
		Codepoints: []int64{10820},
		Characters: "\u2a44",
	},
	"&leftrightvector;": Entity{
		Codepoints: []int64{10574},
		Characters: "\u294e",
	},
	"&lltri;": Entity{
		Codepoints: []int64{9722},
		Characters: "\u25fa",
	},
	"&vsubne;": Entity{
		Codepoints: []int64{10955, 65024},
		Characters: "\u2acb\ufe00",
	},
	"&omid;": Entity{
		Codepoints: []int64{10678},
		Characters: "\u29b6",
	},
	"&vellip;": Entity{
		Codepoints: []int64{8942},
		Characters: "\u22ee",
	},
	"&lne;": Entity{
		Codepoints: []int64{10887},
		Characters: "\u2a87",
	},
	"&acy;": Entity{
		Codepoints: []int64{1072},
		Characters: "\u0430",
	},
	"&xharr;": Entity{
		Codepoints: []int64{10234},
		Characters: "\u27fa",
	},
	"&ltquest;": Entity{
		Codepoints: []int64{10875},
		Characters: "\u2a7b",
	},
	"&ring;": Entity{
		Codepoints: []int64{730},
		Characters: "\u02da",
	},
	"&frac23;": Entity{
		Codepoints: []int64{8532},
		Characters: "\u2154",
	},
	"&frac78;": Entity{
		Codepoints: []int64{8542},
		Characters: "\u215e",
	},
	"&angmsdac;": Entity{
		Codepoints: []int64{10666},
		Characters: "\u29aa",
	},
	"&plusdo;": Entity{
		Codepoints: []int64{8724},
		Characters: "\u2214",
	},
	"&uscr;": Entity{
		Codepoints: []int64{119984},
		Characters: "\U0001d4b0",
	},
	"&tstrok;": Entity{
		Codepoints: []int64{358},
		Characters: "\u0166",
	},
	"&zopf;": Entity{
		Codepoints: []int64{8484},
		Characters: "\u2124",
	},
	"&scirc;": Entity{
		Codepoints: []int64{348},
		Characters: "\u015c",
	},
	"&supplus;": Entity{
		Codepoints: []int64{10944},
		Characters: "\u2ac0",
	},
	"&fscr;": Entity{
		Codepoints: []int64{8497},
		Characters: "\u2131",
	},
	"&nvltrie;": Entity{
		Codepoints: []int64{8884, 8402},
		Characters: "\u22b4\u20d2",
	},
	"&emsp14;": Entity{
		Codepoints: []int64{8197},
		Characters: "\u2005",
	},
	"&tscr;": Entity{
		Codepoints: []int64{119983},
		Characters: "\U0001d4af",
	},
	"&roplus;": Entity{
		Codepoints: []int64{10798},
		Characters: "\u2a2e",
	},
	"&gopf;": Entity{
		Codepoints: []int64{120152},
		Characters: "\U0001d558",
	},
	"&vbarv;": Entity{
		Codepoints: []int64{10985},
		Characters: "\u2ae9",
	},
	"&lcaron;": Entity{
		Codepoints: []int64{318},
		Characters: "\u013e",
	},
	"&nearhk;": Entity{
		Codepoints: []int64{10532},
		Characters: "\u2924",
	},
	"&szlig;": Entity{
		Codepoints: []int64{223},
		Characters: "\u00df",
	},
	"&larrtl;": Entity{
		Codepoints: []int64{8610},
		Characters: "\u21a2",
	},
	"&twixt;": Entity{
		Codepoints: []int64{8812},
		Characters: "\u226c",
	},
	"&emptyverysmallsquare;": Entity{
		Codepoints: []int64{9643},
		Characters: "\u25ab",
	},
	"&topfork;": Entity{
		Codepoints: []int64{10970},
		Characters: "\u2ada",
	},
	"&npolint;": Entity{
		Codepoints: []int64{10772},
		Characters: "\u2a14",
	},
	"&ulcorn;": Entity{
		Codepoints: []int64{8988},
		Characters: "\u231c",
	},
	"&larrfs;": Entity{
		Codepoints: []int64{10525},
		Characters: "\u291d",
	},
	"&langd;": Entity{
		Codepoints: []int64{10641},
		Characters: "\u2991",
	},
	"&ace;": Entity{
		Codepoints: []int64{8766, 819},
		Characters: "\u223e\u0333",
	},
	"&sqcups;": Entity{
		Codepoints: []int64{8852, 65024},
		Characters: "\u2294\ufe00",
	},
	"&smashp;": Entity{
		Codepoints: []int64{10803},
		Characters: "\u2a33",
	},
	"&tcy;": Entity{
		Codepoints: []int64{1058},
		Characters: "\u0422",
	},
	"&wedge;": Entity{
		Codepoints: []int64{8896},
		Characters: "\u22c0",
	},
	"&iprod;": Entity{
		Codepoints: []int64{10812},
		Characters: "\u2a3c",
	},
	"&infintie;": Entity{
		Codepoints: []int64{10717},
		Characters: "\u29dd",
	},
	"&gnsim;": Entity{
		Codepoints: []int64{8935},
		Characters: "\u22e7",
	},
	"&utilde;": Entity{
		Codepoints: []int64{360},
		Characters: "\u0168",
	},
	"&nexist;": Entity{
		Codepoints: []int64{8708},
		Characters: "\u2204",
	},
	"&uarr;": Entity{
		Codepoints: []int64{8657},
		Characters: "\u21d1",
	},
	"&imof;": Entity{
		Codepoints: []int64{8887},
		Characters: "\u22b7",
	},
	"&nu;": Entity{
		Codepoints: []int64{957},
		Characters: "\u03bd",
	},
	"&ecirc;": Entity{
		Codepoints: []int64{234},
		Characters: "\u00ea",
	},
	"&sext;": Entity{
		Codepoints: []int64{10038},
		Characters: "\u2736",
	},
	"&gammad;": Entity{
		Codepoints: []int64{988},
		Characters: "\u03dc",
	},
	"&subdot;": Entity{
		Codepoints: []int64{10941},
		Characters: "\u2abd",
	},
	"&dhar;": Entity{
		Codepoints: []int64{10597},
		Characters: "\u2965",
	},
	"&frac18;": Entity{
		Codepoints: []int64{8539},
		Characters: "\u215b",
	},
	"&cuwed;": Entity{
		Codepoints: []int64{8911},
		Characters: "\u22cf",
	},
	"&boxul;": Entity{
		Codepoints: []int64{9564},
		Characters: "\u255c",
	},
	"&notsquaresubset;": Entity{
		Codepoints: []int64{8847, 824},
		Characters: "\u228f\u0338",
	},
	"&prod;": Entity{
		Codepoints: []int64{8719},
		Characters: "\u220f",
	},
	"&rbrke;": Entity{
		Codepoints: []int64{10636},
		Characters: "\u298c",
	},
	"&gbreve;": Entity{
		Codepoints: []int64{287},
		Characters: "\u011f",
	},
	"&hopf;": Entity{
		Codepoints: []int64{120153},
		Characters: "\U0001d559",
	},
	"&utdot;": Entity{
		Codepoints: []int64{8944},
		Characters: "\u22f0",
	},
	"&simne;": Entity{
		Codepoints: []int64{8774},
		Characters: "\u2246",
	},
	"&csube;": Entity{
		Codepoints: []int64{10961},
		Characters: "\u2ad1",
	},
	"&gcedil;": Entity{
		Codepoints: []int64{290},
		Characters: "\u0122",
	},
	"&planckh;": Entity{
		Codepoints: []int64{8462},
		Characters: "\u210e",
	},
	"&ofcir;": Entity{
		Codepoints: []int64{10687},
		Characters: "\u29bf",
	},
	"&simrarr;": Entity{
		Codepoints: []int64{10610},
		Characters: "\u2972",
	},
	"&ffilig;": Entity{
		Codepoints: []int64{64259},
		Characters: "\ufb03",
	},
	"&nequiv;": Entity{
		Codepoints: []int64{8802},
		Characters: "\u2262",
	},
	"&cedil;": Entity{
		Codepoints: []int64{184},
		Characters: "\u00b8",
	},
	"&napos;": Entity{
		Codepoints: []int64{329},
		Characters: "\u0149",
	},
	"&pr;": Entity{
		Codepoints: []int64{8826},
		Characters: "\u227a",
	},
	"&vert;": Entity{
		Codepoints: []int64{124},
		Characters: "|",
	},
	"&fflig;": Entity{
		Codepoints: []int64{64256},
		Characters: "\ufb00",
	},
	"&erarr;": Entity{
		Codepoints: []int64{10609},
		Characters: "\u2971",
	},
	"&mfr;": Entity{
		Codepoints: []int64{120080},
		Characters: "\U0001d510",
	},
	"&rcub;": Entity{
		Codepoints: []int64{125},
		Characters: "}",
	},
	"&range;": Entity{
		Codepoints: []int64{10661},
		Characters: "\u29a5",
	},
	"&uogon;": Entity{
		Codepoints: []int64{371},
		Characters: "\u0173",
	},
	"&egsdot;": Entity{
		Codepoints: []int64{10904},
		Characters: "\u2a98",
	},
	"&cfr;": Entity{
		Codepoints: []int64{8493},
		Characters: "\u212d",
	},
	"&gcy;": Entity{
		Codepoints: []int64{1075},
		Characters: "\u0433",
	},
	"&rharul;": Entity{
		Codepoints: []int64{10604},
		Characters: "\u296c",
	},
	"&nll;": Entity{
		Codepoints: []int64{8920, 824},
		Characters: "\u22d8\u0338",
	},
	"&drcrop;": Entity{
		Codepoints: []int64{8972},
		Characters: "\u230c",
	},
	"&nesim;": Entity{
		Codepoints: []int64{8770, 824},
		Characters: "\u2242\u0338",
	},
	"&ycirc;": Entity{
		Codepoints: []int64{375},
		Characters: "\u0177",
	},
	"&dot;": Entity{
		Codepoints: []int64{168},
		Characters: "\u00a8",
	},
	"&prcue;": Entity{
		Codepoints: []int64{8828},
		Characters: "\u227c",
	},
	"&ltimes;": Entity{
		Codepoints: []int64{8905},
		Characters: "\u22c9",
	},
	"&rtriltri;": Entity{
		Codepoints: []int64{10702},
		Characters: "\u29ce",
	},
	"&copysr;": Entity{
		Codepoints: []int64{8471},
		Characters: "\u2117",
	},
	"&barvee;": Entity{
		Codepoints: []int64{8893},
		Characters: "\u22bd",
	},
	"&lcy;": Entity{
		Codepoints: []int64{1083},
		Characters: "\u043b",
	},
	"&nvsim;": Entity{
		Codepoints: []int64{8764, 8402},
		Characters: "\u223c\u20d2",
	},
	"&scne;": Entity{
		Codepoints: []int64{10934},
		Characters: "\u2ab6",
	},
	"&ucy;": Entity{
		Codepoints: []int64{1059},
		Characters: "\u0423",
	},
	"&laemptyv;": Entity{
		Codepoints: []int64{10676},
		Characters: "\u29b4",
	},
	"&delta;": Entity{
		Codepoints: []int64{948},
		Characters: "\u03b4",
	},
	"&rpar;": Entity{
		Codepoints: []int64{41},
		Characters: ")",
	},
	"&cap;": Entity{
		Codepoints: []int64{8914},
		Characters: "\u22d2",
	},
	"&notlefttrianglebar;": Entity{
		Codepoints: []int64{10703, 824},
		Characters: "\u29cf\u0338",
	},
	"&vee;": Entity{
		Codepoints: []int64{8897},
		Characters: "\u22c1",
	},
	"&angmsdad;": Entity{
		Codepoints: []int64{10667},
		Characters: "\u29ab",
	},
	"&kappa;": Entity{
		Codepoints: []int64{954},
		Characters: "\u03ba",
	},
	"&downrightvectorbar;": Entity{
		Codepoints: []int64{10583},
		Characters: "\u2957",
	},
	"&sdote;": Entity{
		Codepoints: []int64{10854},
		Characters: "\u2a66",
	},
	"&gvne;": Entity{
		Codepoints: []int64{8809, 65024},
		Characters: "\u2269\ufe00",
	},
	"&not;": Entity{
		Codepoints: []int64{172},
		Characters: "\u00ac",
	},
	"&nap;": Entity{
		Codepoints: []int64{8777},
		Characters: "\u2249",
	},
	"&excl;": Entity{
		Codepoints: []int64{33},
		Characters: "!",
	},
	"&ngtv;": Entity{
		Codepoints: []int64{8811, 824},
		Characters: "\u226b\u0338",
	},
	"&greatergreater;": Entity{
		Codepoints: []int64{10914},
		Characters: "\u2aa2",
	},
	"&copf;": Entity{
		Codepoints: []int64{120148},
		Characters: "\U0001d554",
	},
	"&parsl;": Entity{
		Codepoints: []int64{11005},
		Characters: "\u2afd",
	},
	"&xlarr;": Entity{
		Codepoints: []int64{10229},
		Characters: "\u27f5",
	},
	"&nrarrw;": Entity{
		Codepoints: []int64{8605, 824},
		Characters: "\u219d\u0338",
	},
	"&harrcir;": Entity{
		Codepoints: []int64{10568},
		Characters: "\u2948",
	},
	"&andd;": Entity{
		Codepoints: []int64{10844},
		Characters: "\u2a5c",
	},
	"&bsim;": Entity{
		Codepoints: []int64{8765},
		Characters: "\u223d",
	},
	"&natur;": Entity{
		Codepoints: []int64{9838},
		Characters: "\u266e",
	},
	"&hstrok;": Entity{
		Codepoints: []int64{294},
		Characters: "\u0126",
	},
	"&forkv;": Entity{
		Codepoints: []int64{10969},
		Characters: "\u2ad9",
	},
	"&ccaps;": Entity{
		Codepoints: []int64{10829},
		Characters: "\u2a4d",
	},
	"&gtdot;": Entity{
		Codepoints: []int64{8919},
		Characters: "\u22d7",
	},
	"&zcy;": Entity{
		Codepoints: []int64{1047},
		Characters: "\u0417",
	},
	"&uparrowbar;": Entity{
		Codepoints: []int64{10514},
		Characters: "\u2912",
	},
	"&bcong;": Entity{
		Codepoints: []int64{8780},
		Characters: "\u224c",
	},
	"&lpar;": Entity{
		Codepoints: []int64{40},
		Characters: "(",
	},
	"&wr;": Entity{
		Codepoints: []int64{8768},
		Characters: "\u2240",
	},
	"&lfloor;": Entity{
		Codepoints: []int64{8970},
		Characters: "\u230a",
	},
	"&phiv;": Entity{
		Codepoints: []int64{981},
		Characters: "\u03d5",
	},
	"&npr;": Entity{
		Codepoints: []int64{8832},
		Characters: "\u2280",
	},
	"&rdldhar;": Entity{
		Codepoints: []int64{10601},
		Characters: "\u2969",
	},
	"&ogon;": Entity{
		Codepoints: []int64{731},
		Characters: "\u02db",
	},
	"&kgreen;": Entity{
		Codepoints: []int64{312},
		Characters: "\u0138",
	},
	"&amalg;": Entity{
		Codepoints: []int64{10815},
		Characters: "\u2a3f",
	},
	"&numsp;": Entity{
		Codepoints: []int64{8199},
		Characters: "\u2007",
	},
	"&dzcy;": Entity{
		Codepoints: []int64{1039},
		Characters: "\u040f",
	},
	"&mdash;": Entity{
		Codepoints: []int64{8212},
		Characters: "\u2014",
	},
	"&forall;": Entity{
		Codepoints: []int64{8704},
		Characters: "\u2200",
	},
	"&thinsp;": Entity{
		Codepoints: []int64{8201},
		Characters: "\u2009",
	},
	"&lbarr;": Entity{
		Codepoints: []int64{10508},
		Characters: "\u290c",
	},
	"&shy;": Entity{
		Codepoints: []int64{173},
		Characters: "\u00ad",
	},
	"&jcy;": Entity{
		Codepoints: []int64{1049},
		Characters: "\u0419",
	},
	"&dfr;": Entity{
		Codepoints: []int64{120071},
		Characters: "\U0001d507",
	},
	"&jcirc;": Entity{
		Codepoints: []int64{309},
		Characters: "\u0135",
	},
	"&gesdoto;": Entity{
		Codepoints: []int64{10882},
		Characters: "\u2a82",
	},
	"&idot;": Entity{
		Codepoints: []int64{304},
		Characters: "\u0130",
	},
	"&nldr;": Entity{
		Codepoints: []int64{8229},
		Characters: "\u2025",
	},
	"&suplarr;": Entity{
		Codepoints: []int64{10619},
		Characters: "\u297b",
	},
	"&vzigzag;": Entity{
		Codepoints: []int64{10650},
		Characters: "\u299a",
	},
	"&bbrktbrk;": Entity{
		Codepoints: []int64{9142},
		Characters: "\u23b6",
	},
	"&efdot;": Entity{
		Codepoints: []int64{8786},
		Characters: "\u2252",
	},
	"&notin;": Entity{
		Codepoints: []int64{8713},
		Characters: "\u2209",
	},
	"&igrave;": Entity{
		Codepoints: []int64{236},
		Characters: "\u00ec",
	},
	"&searhk;": Entity{
		Codepoints: []int64{10533},
		Characters: "\u2925",
	},
	"&nsube;": Entity{
		Codepoints: []int64{10949, 824},
		Characters: "\u2ac5\u0338",
	},
	"&prsim;": Entity{
		Codepoints: []int64{8830},
		Characters: "\u227e",
	},
	"&dharr;": Entity{
		Codepoints: []int64{8642},
		Characters: "\u21c2",
	},
	"&lesg;": Entity{
		Codepoints: []int64{8922, 65024},
		Characters: "\u22da\ufe00",
	},
	"&bnot;": Entity{
		Codepoints: []int64{8976},
		Characters: "\u2310",
	},
	"&pointint;": Entity{
		Codepoints: []int64{10773},
		Characters: "\u2a15",
	},
	"&rsh;": Entity{
		Codepoints: []int64{8625},
		Characters: "\u21b1",
	},
	"&marker;": Entity{
		Codepoints: []int64{9646},
		Characters: "\u25ae",
	},
	"&eth;": Entity{
		Codepoints: []int64{240},
		Characters: "\u00f0",
	},
	"&ccedil;": Entity{
		Codepoints: []int64{231},
		Characters: "\u00e7",
	},
	"&setmn;": Entity{
		Codepoints: []int64{8726},
		Characters: "\u2216",
	},
	"&mapstoleft;": Entity{
		Codepoints: []int64{8612},
		Characters: "\u21a4",
	},
	"&gel;": Entity{
		Codepoints: []int64{10892},
		Characters: "\u2a8c",
	},
	"&raemptyv;": Entity{
		Codepoints: []int64{10675},
		Characters: "\u29b3",
	},
	"&xcirc;": Entity{
		Codepoints: []int64{9711},
		Characters: "\u25ef",
	},
	"&utri;": Entity{
		Codepoints: []int64{9653},
		Characters: "\u25b5",
	},
	"&isinsv;": Entity{
		Codepoints: []int64{8947},
		Characters: "\u22f3",
	},
	"&rarrc;": Entity{
		Codepoints: []int64{10547},
		Characters: "\u2933",
	},
	"&lefttrianglebar;": Entity{
		Codepoints: []int64{10703},
		Characters: "\u29cf",
	},
	"&dzigrarr;": Entity{
		Codepoints: []int64{10239},
		Characters: "\u27ff",
	},
	"&lobrk;": Entity{
		Codepoints: []int64{10214},
		Characters: "\u27e6",
	},
	"&pcy;": Entity{
		Codepoints: []int64{1055},
		Characters: "\u041f",
	},
	"&topbot;": Entity{
		Codepoints: []int64{9014},
		Characters: "\u2336",
	},
	"&there4;": Entity{
		Codepoints: []int64{8756},
		Characters: "\u2234",
	},
	"&leftupteevector;": Entity{
		Codepoints: []int64{10592},
		Characters: "\u2960",
	},
	"&becaus;": Entity{
		Codepoints: []int64{8757},
		Characters: "\u2235",
	},
	"&vdash;": Entity{
		Codepoints: []int64{8873},
		Characters: "\u22a9",
	},
	"&larrlp;": Entity{
		Codepoints: []int64{8619},
		Characters: "\u21ab",
	},
	"&boxvr;": Entity{
		Codepoints: []int64{9567},
		Characters: "\u255f",
	},
	"&simle;": Entity{
		Codepoints: []int64{10911},
		Characters: "\u2a9f",
	},
	"&lozf;": Entity{
		Codepoints: []int64{10731},
		Characters: "\u29eb",
	},
	"&nlarr;": Entity{
		Codepoints: []int64{8653},
		Characters: "\u21cd",
	},
	"&uring;": Entity{
		Codepoints: []int64{367},
		Characters: "\u016f",
	},
	"&gtcir;": Entity{
		Codepoints: []int64{10874},
		Characters: "\u2a7a",
	},
	"&tscy;": Entity{
		Codepoints: []int64{1094},
		Characters: "\u0446",
	},
	"&cdot;": Entity{
		Codepoints: []int64{266},
		Characters: "\u010a",
	},
	"&gimel;": Entity{
		Codepoints: []int64{8503},
		Characters: "\u2137",
	},
	"&hfr;": Entity{
		Codepoints: []int64{8460},
		Characters: "\u210c",
	},
	"&gap;": Entity{
		Codepoints: []int64{10886},
		Characters: "\u2a86",
	},
	"&ldsh;": Entity{
		Codepoints: []int64{8626},
		Characters: "\u21b2",
	},
	"&jfr;": Entity{
		Codepoints: []int64{120077},
		Characters: "\U0001d50d",
	},
	"&map;": Entity{
		Codepoints: []int64{8614},
		Characters: "\u21a6",
	},
	"&mldr;": Entity{
		Codepoints: []int64{8230},
		Characters: "\u2026",
	},
	"&nbumpe;": Entity{
		Codepoints: []int64{8783, 824},
		Characters: "\u224f\u0338",
	},
	"&racute;": Entity{
		Codepoints: []int64{341},
		Characters: "\u0155",
	},
	"&kscr;": Entity{
		Codepoints: []int64{119974},
		Characters: "\U0001d4a6",
	},
	"&eqvparsl;": Entity{
		Codepoints: []int64{10725},
		Characters: "\u29e5",
	},
	"&lowbar;": Entity{
		Codepoints: []int64{95},
		Characters: "_",
	},
	"&subne;": Entity{
		Codepoints: []int64{8842},
		Characters: "\u228a",
	},
	"&pscr;": Entity{
		Codepoints: []int64{119979},
		Characters: "\U0001d4ab",
	},
	"&lbbrk;": Entity{
		Codepoints: []int64{10098},
		Characters: "\u2772",
	},
	"&late;": Entity{
		Codepoints: []int64{10925},
		Characters: "\u2aad",
	},
	"&gscr;": Entity{
		Codepoints: []int64{119970},
		Characters: "\U0001d4a2",
	},
	"&lrhar;": Entity{
		Codepoints: []int64{8651},
		Characters: "\u21cb",
	},
	"&bscr;": Entity{
		Codepoints: []int64{119991},
		Characters: "\U0001d4b7",
	},
	"&udarr;": Entity{
		Codepoints: []int64{8645},
		Characters: "\u21c5",
	},
	"&ldrdhar;": Entity{
		Codepoints: []int64{10599},
		Characters: "\u2967",
	},
	"&lat;": Entity{
		Codepoints: []int64{10923},
		Characters: "\u2aab",
	},
	"&umacr;": Entity{
		Codepoints: []int64{362},
		Characters: "\u016a",
	},
	"&rsaquo;": Entity{
		Codepoints: []int64{8250},
		Characters: "\u203a",
	},
	"&harrw;": Entity{
		Codepoints: []int64{8621},
		Characters: "\u21ad",
	},
	"&gesdot;": Entity{
		Codepoints: []int64{10880},
		Characters: "\u2a80",
	},
	"&rightupdownvector;": Entity{
		Codepoints: []int64{10575},
		Characters: "\u294f",
	},
	"&cscr;": Entity{
		Codepoints: []int64{119992},
		Characters: "\U0001d4b8",
	},
	"&bot;": Entity{
		Codepoints: []int64{8869},
		Characters: "\u22a5",
	},
	"&gt;": Entity{
		Codepoints: []int64{62},
		Characters: ">",
	},
	"&int;": Entity{
		Codepoints: []int64{8747},
		Characters: "\u222b",
	},
	"&nopf;": Entity{
		Codepoints: []int64{8469},
		Characters: "\u2115",
	},
	"&sum;": Entity{
		Codepoints: []int64{8721},
		Characters: "\u2211",
	},
	"&qint;": Entity{
		Codepoints: []int64{10764},
		Characters: "\u2a0c",
	},
	"&tprime;": Entity{
		Codepoints: []int64{8244},
		Characters: "\u2034",
	},
	"&subsup;": Entity{
		Codepoints: []int64{10963},
		Characters: "\u2ad3",
	},
	"&smt;": Entity{
		Codepoints: []int64{10922},
		Characters: "\u2aaa",
	},
	"&underparenthesis;": Entity{
		Codepoints: []int64{9181},
		Characters: "\u23dd",
	},
	"&prnap;": Entity{
		Codepoints: []int64{10937},
		Characters: "\u2ab9",
	},
	"&notnestedlessless;": Entity{
		Codepoints: []int64{10913, 824},
		Characters: "\u2aa1\u0338",
	},
	"&dscr;": Entity{
		Codepoints: []int64{119993},
		Characters: "\U0001d4b9",
	},
	"&nharr;": Entity{
		Codepoints: []int64{8622},
		Characters: "\u21ae",
	},
	"&subrarr;": Entity{
		Codepoints: []int64{10617},
		Characters: "\u2979",
	},
	"&afr;": Entity{
		Codepoints: []int64{120094},
		Characters: "\U0001d51e",
	},
	"&llhard;": Entity{
		Codepoints: []int64{10603},
		Characters: "\u296b",
	},
	"&nsime;": Entity{
		Codepoints: []int64{8772},
		Characters: "\u2244",
	},
	"&yucy;": Entity{
		Codepoints: []int64{1102},
		Characters: "\u044e",
	},
	"&ddotrahd;": Entity{
		Codepoints: []int64{10513},
		Characters: "\u2911",
	},
	"&iinfin;": Entity{
		Codepoints: []int64{10716},
		Characters: "\u29dc",
	},
	"&odiv;": Entity{
		Codepoints: []int64{10808},
		Characters: "\u2a38",
	},
	"&leftdownvectorbar;": Entity{
		Codepoints: []int64{10585},
		Characters: "\u2959",
	},
	"&rtrie;": Entity{
		Codepoints: []int64{8885},
		Characters: "\u22b5",
	},
	"&euml;": Entity{
		Codepoints: []int64{203},
		Characters: "\u00cb",
	},
	"&tdot;": Entity{
		Codepoints: []int64{8411},
		Characters: "\u20db",
	},
	"&gescc;": Entity{
		Codepoints: []int64{10921},
		Characters: "\u2aa9",
	},
	"&oplus;": Entity{
		Codepoints: []int64{8853},
		Characters: "\u2295",
	},
	"&frac14;": Entity{
		Codepoints: []int64{188},
		Characters: "\u00bc",
	},
	"&apid;": Entity{
		Codepoints: []int64{8779},
		Characters: "\u224b",
	},
	"&npart;": Entity{
		Codepoints: []int64{8706, 824},
		Characters: "\u2202\u0338",
	},
	"&phone;": Entity{
		Codepoints: []int64{9742},
		Characters: "\u260e",
	},
	"&righttrianglebar;": Entity{
		Codepoints: []int64{10704},
		Characters: "\u29d0",
	},
	"&plusdu;": Entity{
		Codepoints: []int64{10789},
		Characters: "\u2a25",
	},
	"&cross;": Entity{
		Codepoints: []int64{10007},
		Characters: "\u2717",
	},
	"&nisd;": Entity{
		Codepoints: []int64{8954},
		Characters: "\u22fa",
	},
	"&zwnj;": Entity{
		Codepoints: []int64{8204},
		Characters: "\u200c",
	},
	"&aelig;": Entity{
		Codepoints: []int64{198},
		Characters: "\u00c6",
	},
	"&hoarr;": Entity{
		Codepoints: []int64{8703},
		Characters: "\u21ff",
	},
	"&mcomma;": Entity{
		Codepoints: []int64{10793},
		Characters: "\u2a29",
	},
	"&ecolon;": Entity{
		Codepoints: []int64{8789},
		Characters: "\u2255",
	},
	"&nacute;": Entity{
		Codepoints: []int64{323},
		Characters: "\u0143",
	},
	"&sacute;": Entity{
		Codepoints: []int64{346},
		Characters: "\u015a",
	},
	"&eopf;": Entity{
		Codepoints: []int64{120150},
		Characters: "\U0001d556",
	},
	"&isine;": Entity{
		Codepoints: []int64{8953},
		Characters: "\u22f9",
	},
	"&lbrke;": Entity{
		Codepoints: []int64{10635},
		Characters: "\u298b",
	},
	"&rotimes;": Entity{
		Codepoints: []int64{10805},
		Characters: "\u2a35",
	},
	"&notnivb;": Entity{
		Codepoints: []int64{8958},
		Characters: "\u22fe",
	},
	"&rfloor;": Entity{
		Codepoints: []int64{8971},
		Characters: "\u230b",
	},
	"&bowtie;": Entity{
		Codepoints: []int64{8904},
		Characters: "\u22c8",
	},
	"&af;": Entity{
		Codepoints: []int64{8289},
		Characters: "\u2061",
	},
	"&rect;": Entity{
		Codepoints: []int64{9645},
		Characters: "\u25ad",
	},
	"&gnap;": Entity{
		Codepoints: []int64{10890},
		Characters: "\u2a8a",
	},
	"&overparenthesis;": Entity{
		Codepoints: []int64{9180},
		Characters: "\u23dc",
	},
	"&ffllig;": Entity{
		Codepoints: []int64{64260},
		Characters: "\ufb04",
	},
	"&lambda;": Entity{
		Codepoints: []int64{923},
		Characters: "\u039b",
	},
	"&num;": Entity{
		Codepoints: []int64{35},
		Characters: "#",
	},
	"&vopf;": Entity{
		Codepoints: []int64{120141},
		Characters: "\U0001d54d",
	},
	"&starf;": Entity{
		Codepoints: []int64{9733},
		Characters: "\u2605",
	},
	"&telrec;": Entity{
		Codepoints: []int64{8981},
		Characters: "\u2315",
	},
	"&wfr;": Entity{
		Codepoints: []int64{120116},
		Characters: "\U0001d534",
	},
	"&larrhk;": Entity{
		Codepoints: []int64{8617},
		Characters: "\u21a9",
	},
	"&frac58;": Entity{
		Codepoints: []int64{8541},
		Characters: "\u215d",
	},
	"&xnis;": Entity{
		Codepoints: []int64{8955},
		Characters: "\u22fb",
	},
	"&filledsmallsquare;": Entity{
		Codepoints: []int64{9724},
		Characters: "\u25fc",
	},
	"&kappav;": Entity{
		Codepoints: []int64{1008},
		Characters: "\u03f0",
	},
	"&notinvb;": Entity{
		Codepoints: []int64{8951},
		Characters: "\u22f7",
	},
	"&dollar;": Entity{
		Codepoints: []int64{36},
		Characters: "$",
	},
	"&cire;": Entity{
		Codepoints: []int64{10691},
		Characters: "\u29c3",
	},
	"&iacute;": Entity{
		Codepoints: []int64{237},
		Characters: "\u00ed",
	},
	"&vdashl;": Entity{
		Codepoints: []int64{10982},
		Characters: "\u2ae6",
	},
	"&ocirc;": Entity{
		Codepoints: []int64{212},
		Characters: "\u00d4",
	},
	"&boxhu;": Entity{
		Codepoints: []int64{9576},
		Characters: "\u2568",
	},
	"&cir;": Entity{
		Codepoints: []int64{9675},
		Characters: "\u25cb",
	},
	"&ncy;": Entity{
		Codepoints: []int64{1085},
		Characters: "\u043d",
	},
	"&rtri;": Entity{
		Codepoints: []int64{9657},
		Characters: "\u25b9",
	},
	"&ccups;": Entity{
		Codepoints: []int64{10828},
		Characters: "\u2a4c",
	},
	"&rightvectorbar;": Entity{
		Codepoints: []int64{10579},
		Characters: "\u2953",
	},
	"&lhblk;": Entity{
		Codepoints: []int64{9604},
		Characters: "\u2584",
	},
	"&wp;": Entity{
		Codepoints: []int64{8472},
		Characters: "\u2118",
	},
	"&rarrap;": Entity{
		Codepoints: []int64{10613},
		Characters: "\u2975",
	},
	"&zdot;": Entity{
		Codepoints: []int64{379},
		Characters: "\u017b",
	},
	"&abreve;": Entity{
		Codepoints: []int64{259},
		Characters: "\u0103",
	},
	"&ovbar;": Entity{
		Codepoints: []int64{9021},
		Characters: "\u233d",
	},
	"&par;": Entity{
		Codepoints: []int64{8741},
		Characters: "\u2225",
	},
	"&nobreak;": Entity{
		Codepoints: []int64{8288},
		Characters: "\u2060",
	},
	"&aacute;": Entity{
		Codepoints: []int64{193},
		Characters: "\u00c1",
	},
	"&rang;": Entity{
		Codepoints: []int64{10217},
		Characters: "\u27e9",
	},
	"&ltdot;": Entity{
		Codepoints: []int64{8918},
		Characters: "\u22d6",
	},
	"&period;": Entity{
		Codepoints: []int64{46},
		Characters: ".",
	},
	"&or;": Entity{
		Codepoints: []int64{10836},
		Characters: "\u2a54",
	},
	"&rtrif;": Entity{
		Codepoints: []int64{9656},
		Characters: "\u25b8",
	},
	"&rarrsim;": Entity{
		Codepoints: []int64{10612},
		Characters: "\u2974",
	},
	"&prne;": Entity{
		Codepoints: []int64{10933},
		Characters: "\u2ab5",
	},
	"&hscr;": Entity{
		Codepoints: []int64{8459},
		Characters: "\u210b",
	},
	"&lharul;": Entity{
		Codepoints: []int64{10602},
		Characters: "\u296a",
	},
	"&rppolint;": Entity{
		Codepoints: []int64{10770},
		Characters: "\u2a12",
	},
	"&capcup;": Entity{
		Codepoints: []int64{10823},
		Characters: "\u2a47",
	},
	"&tosa;": Entity{
		Codepoints: []int64{10537},
		Characters: "\u2929",
	},
	"&angmsdah;": Entity{
		Codepoints: []int64{10671},
		Characters: "\u29af",
	},
	"&iocy;": Entity{
		Codepoints: []int64{1025},
		Characters: "\u0401",
	},
	"&pluscir;": Entity{
		Codepoints: []int64{10786},
		Characters: "\u2a22",
	},
	"&ohbar;": Entity{
		Codepoints: []int64{10677},
		Characters: "\u29b5",
	},
	"&fltns;": Entity{
		Codepoints: []int64{9649},
		Characters: "\u25b1",
	},
	"&popf;": Entity{
		Codepoints: []int64{120161},
		Characters: "\U0001d561",
	},
	"&assign;": Entity{
		Codepoints: []int64{8788},
		Characters: "\u2254",
	},
	"&ctdot;": Entity{
		Codepoints: []int64{8943},
		Characters: "\u22ef",
	},
	"&sqsube;": Entity{
		Codepoints: []int64{8849},
		Characters: "\u2291",
	},
	"&nsce;": Entity{
		Codepoints: []int64{10928, 824},
		Characters: "\u2ab0\u0338",
	},
	"&bsime;": Entity{
		Codepoints: []int64{8909},
		Characters: "\u22cd",
	},
	"&oelig;": Entity{
		Codepoints: []int64{339},
		Characters: "\u0153",
	},
	"&leftdownteevector;": Entity{
		Codepoints: []int64{10593},
		Characters: "\u2961",
	},
	"&chi;": Entity{
		Codepoints: []int64{935},
		Characters: "\u03a7",
	},
	"&acd;": Entity{
		Codepoints: []int64{8767},
		Characters: "\u223f",
	},
	"&psi;": Entity{
		Codepoints: []int64{936},
		Characters: "\u03a8",
	},
	"&trie;": Entity{
		Codepoints: []int64{8796},
		Characters: "\u225c",
	},
	"&topcir;": Entity{
		Codepoints: []int64{10993},
		Characters: "\u2af1",
	},
	"&nge;": Entity{
		Codepoints: []int64{8807, 824},
		Characters: "\u2267\u0338",
	},
	"&nhpar;": Entity{
		Codepoints: []int64{10994},
		Characters: "\u2af2",
	},
	"&nsim;": Entity{
		Codepoints: []int64{8769},
		Characters: "\u2241",
	},
	"&orarr;": Entity{
		Codepoints: []int64{8635},
		Characters: "\u21bb",
	},
	"&ltri;": Entity{
		Codepoints: []int64{9667},
		Characters: "\u25c3",
	},
	"&downleftvectorbar;": Entity{
		Codepoints: []int64{10582},
		Characters: "\u2956",
	},
	"&jukcy;": Entity{
		Codepoints: []int64{1108},
		Characters: "\u0454",
	},
	"&apacir;": Entity{
		Codepoints: []int64{10863},
		Characters: "\u2a6f",
	},
	"&intcal;": Entity{
		Codepoints: []int64{8890},
		Characters: "\u22ba",
	},
	"&lsh;": Entity{
		Codepoints: []int64{8624},
		Characters: "\u21b0",
	},
	"&llarr;": Entity{
		Codepoints: []int64{8647},
		Characters: "\u21c7",
	},
	"&qprime;": Entity{
		Codepoints: []int64{8279},
		Characters: "\u2057",
	},
	"&os;": Entity{
		Codepoints: []int64{9416},
		Characters: "\u24c8",
	},
	"&race;": Entity{
		Codepoints: []int64{8765, 817},
		Characters: "\u223d\u0331",
	},
	"&njcy;": Entity{
		Codepoints: []int64{1114},
		Characters: "\u045a",
	},
	"&lsqb;": Entity{
		Codepoints: []int64{91},
		Characters: "[",
	},
	"&rbrkslu;": Entity{
		Codepoints: []int64{10640},
		Characters: "\u2990",
	},
	"&jopf;": Entity{
		Codepoints: []int64{120155},
		Characters: "\U0001d55b",
	},
	"&comma;": Entity{
		Codepoints: []int64{44},
		Characters: ",",
	},
	"&zwj;": Entity{
		Codepoints: []int64{8205},
		Characters: "\u200d",
	},
	"&pi;": Entity{
		Codepoints: []int64{960},
		Characters: "\u03c0",
	},
	"&agrave;": Entity{
		Codepoints: []int64{192},
		Characters: "\u00c0",
	},
	"&icirc;": Entity{
		Codepoints: []int64{238},
		Characters: "\u00ee",
	},
	"&tcaron;": Entity{
		Codepoints: []int64{357},
		Characters: "\u0165",
	},
	"&djcy;": Entity{
		Codepoints: []int64{1026},
		Characters: "\u0402",
	},
	"&xopf;": Entity{
		Codepoints: []int64{120143},
		Characters: "\U0001d54f",
	},
	"&atilde;": Entity{
		Codepoints: []int64{195},
		Characters: "\u00c3",
	},
	"&boxhd;": Entity{
		Codepoints: []int64{9573},
		Characters: "\u2565",
	},
	"&hearts;": Entity{
		Codepoints: []int64{9829},
		Characters: "\u2665",
	},
	"&supdsub;": Entity{
		Codepoints: []int64{10968},
		Characters: "\u2ad8",
	},
	"&downleftteevector;": Entity{
		Codepoints: []int64{10590},
		Characters: "\u295e",
	},
	"&auml;": Entity{
		Codepoints: []int64{228},
		Characters: "\u00e4",
	},
	"&ijlig;": Entity{
		Codepoints: []int64{307},
		Characters: "\u0133",
	},
	"&ecir;": Entity{
		Codepoints: []int64{8790},
		Characters: "\u2256",
	},
	"&ratio;": Entity{
		Codepoints: []int64{8758},
		Characters: "\u2236",
	},
	"&vangrt;": Entity{
		Codepoints: []int64{10652},
		Characters: "\u299c",
	},
	"&softcy;": Entity{
		Codepoints: []int64{1068},
		Characters: "\u042c",
	},
	"&im;": Entity{
		Codepoints: []int64{8465},
		Characters: "\u2111",
	},
	"&reg;": Entity{
		Codepoints: []int64{174},
		Characters: "\u00ae",
	},
	"&hairsp;": Entity{
		Codepoints: []int64{8202},
		Characters: "\u200a",
	},
	"&swarhk;": Entity{
		Codepoints: []int64{10534},
		Characters: "\u2926",
	},
	"&emacr;": Entity{
		Codepoints: []int64{275},
		Characters: "\u0113",
	},
	"&escr;": Entity{
		Codepoints: []int64{8495},
		Characters: "\u212f",
	},
	"&nsc;": Entity{
		Codepoints: []int64{8833},
		Characters: "\u2281",
	},
	"&sqsub;": Entity{
		Codepoints: []int64{8847},
		Characters: "\u228f",
	},
	"&wcirc;": Entity{
		Codepoints: []int64{373},
		Characters: "\u0175",
	},
	"&plussim;": Entity{
		Codepoints: []int64{10790},
		Characters: "\u2a26",
	},
	"&demptyv;": Entity{
		Codepoints: []int64{10673},
		Characters: "\u29b1",
	},
	"&rceil;": Entity{
		Codepoints: []int64{8969},
		Characters: "\u2309",
	},
	"&lsaquo;": Entity{
		Codepoints: []int64{8249},
		Characters: "\u2039",
	},
	"&tcedil;": Entity{
		Codepoints: []int64{355},
		Characters: "\u0163",
	},
	"&frac13;": Entity{
		Codepoints: []int64{8531},
		Characters: "\u2153",
	},
	"&trade;": Entity{
		Codepoints: []int64{8482},
		Characters: "\u2122",
	},
	"&nwarr;": Entity{
		Codepoints: []int64{8598},
		Characters: "\u2196",
	},
	"&andand;": Entity{
		Codepoints: []int64{10837},
		Characters: "\u2a55",
	},
	"&rhard;": Entity{
		Codepoints: []int64{8641},
		Characters: "\u21c1",
	},
	"&lbrkslu;": Entity{
		Codepoints: []int64{10637},
		Characters: "\u298d",
	},
	"&subplus;": Entity{
		Codepoints: []int64{10943},
		Characters: "\u2abf",
	},
	"&dfisht;": Entity{
		Codepoints: []int64{10623},
		Characters: "\u297f",
	},
	"&lesdoto;": Entity{
		Codepoints: []int64{10881},
		Characters: "\u2a81",
	},
	"&sup2;": Entity{
		Codepoints: []int64{178},
		Characters: "\u00b2",
	},
	"&amp;": Entity{
		Codepoints: []int64{38},
		Characters: "&",
	},
	"&hybull;": Entity{
		Codepoints: []int64{8259},
		Characters: "\u2043",
	},
	"&ohm;": Entity{
		Codepoints: []int64{937},
		Characters: "\u03a9",
	},
	"&kcy;": Entity{
		Codepoints: []int64{1082},
		Characters: "\u043a",
	},
	"&nrtrie;": Entity{
		Codepoints: []int64{8941},
		Characters: "\u22ed",
	},
	"&prime;": Entity{
		Codepoints: []int64{8243},
		Characters: "\u2033",
	},
	"&zscr;": Entity{
		Codepoints: []int64{119989},
		Characters: "\U0001d4b5",
	},
	"&para;": Entity{
		Codepoints: []int64{182},
		Characters: "\u00b6",
	},
	"&epsi;": Entity{
		Codepoints: []int64{949},
		Characters: "\u03b5",
	},
	"&fpartint;": Entity{
		Codepoints: []int64{10765},
		Characters: "\u2a0d",
	},
	"&smile;": Entity{
		Codepoints: []int64{8995},
		Characters: "\u2323",
	},
	"&orslope;": Entity{
		Codepoints: []int64{10839},
		Characters: "\u2a57",
	},
	"&bull;": Entity{
		Codepoints: []int64{8226},
		Characters: "\u2022",
	},
	"&operp;": Entity{
		Codepoints: []int64{10681},
		Characters: "\u29b9",
	},
	"&ocy;": Entity{
		Codepoints: []int64{1054},
		Characters: "\u041e",
	},
	"&lstrok;": Entity{
		Codepoints: []int64{321},
		Characters: "\u0141",
	},
	"&swnwar;": Entity{
		Codepoints: []int64{10538},
		Characters: "\u292a",
	},
	"&ensp;": Entity{
		Codepoints: []int64{8194},
		Characters: "\u2002",
	},
	"&check;": Entity{
		Codepoints: []int64{10003},
		Characters: "\u2713",
	},
	"&malt;": Entity{
		Codepoints: []int64{10016},
		Characters: "\u2720",
	},
	"&lfisht;": Entity{
		Codepoints: []int64{10620},
		Characters: "\u297c",
	},
}
