package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type Trace struct {
	TraceArray []ObjTrace `json:"trace"`
}

type S3Class struct {
	name string
	storagePrice []float64
	getPrice float64
	postPrice float64
	retrievalPrice []float64
	retrievalData []float64
	days int
}

type ObjClass struct {
	className string
	days int
	daysCountingOntBill int
	BillingDays int
	cost float64
	size int
	getRequests float64
	postRequests float64
	retrievalRequests []float64

}
var classes = []S3Class {
	{name:"archive", storagePrice: []float64{0.00099}, getPrice: 0.0000004, postPrice: 0.00005, retrievalPrice: []float64{0,0.00010, 0.000025}, retrievalData: []float64{0, 0.02, 0.0025}, days: 180},
	{name:"glacier", storagePrice: []float64{0.004}, getPrice: 0.0000004, postPrice: 0.00005, retrievalPrice: []float64{0.01, 0.00005, 0.000025}, retrievalData: []float64{0.03, 0.01, 0.0025}, days: 90},
	{name:"infrequent access", storagePrice: []float64{0.0125}, getPrice: 0.001, postPrice: 0.01, retrievalData: []float64{0.01}, days: 30},
	{name:"standard", storagePrice: []float64{0.023, 0.022, 0.021}, getPrice: 0.0000004, postPrice: 0.000005, days: 30},

}

//corresponds to the trace of an S3 object
type ObjTrace struct {
	ObjSize   int `json:"objSize"`
	Requests map[int][]int `json:"requests"`
}

func main() {
	f, err := os.Create("/output_results")
	defer f.Close()
	w := bufio.NewWriter(f)


	//fileName = "input_test.json"
	//fileName := "/home/livi/Desktop/simul_tcc_py/few_access.json"
	//fileName := "/home/livi/easyTimes100.json"
	fileName:= "/home/livi/Desktop/simul_tcc_py/easyTimes10000.json"
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var traces Trace

	json.Unmarshal(byteValue, &traces)
	sizeEnd := 1000
	for trace:=range traces.TraceArray {
		fmt.Println(trace)
		t := ObjTrace{
			ObjSize:  6000,
			Requests: map[int][]int{1: []int{0, 0}, 2: []int{0, 0}, 3: []int{0, 0}, 4: []int{0, 0}, 5: []int{0, 0}, 6: []int{0, 0}, 7: []int{0, 0}, 8: []int{0, 0}, 9: []int{0, 0}, 10: []int{0, 0}, 11: []int{0, 0}, 12: []int{0, 0}, 13: []int{0, 0}, 14: []int{0, 0}, 15: []int{0, 0}, 16: []int{0, 0}, 17: []int{0, 0}, 18: []int{0, 0}, 19: []int{0, 0}, 20: []int{0, 0}, 21: []int{0, 0}, 22: []int{0, 0}, 23: []int{0, 0}, 24: []int{0, 0}, 25: []int{0, 0}, 26: []int{0, 0}, 27: []int{0, 0}, 28: []int{0, 0}, 29: []int{0, 0}, 30: []int{0, 0}, 31: []int{0, 0}, 32: []int{0, 0}, 33: []int{0, 0}, 34: []int{0, 0}, 35: []int{0, 0}, 36: []int{2490000, 1040000}, 37: []int{540000, 940000}, 38: []int{2260000, 3640000}, 39: []int{0, 0}, 40: []int{840000, 3610000}, 41: []int{2350000, 3250000}, 42: []int{1030000, 4530000}, 43: []int{4130000, 2460000}, 44: []int{3850000, 3490000}, 45: []int{1990000, 910000}, 46: []int{0, 0}, 47: []int{0, 0}, 48: []int{0, 0}, 49: []int{0, 0}, 50: []int{0, 0}, 51: []int{0, 0}, 52: []int{0, 0}, 53: []int{0, 0}, 54: []int{0, 0}, 55: []int{0, 0}, 56: []int{4280000, 2410000}, 57: []int{0, 0}, 58: []int{0, 0}, 59: []int{0, 0}, 60: []int{0, 0}, 61: []int{0, 0}, 62: []int{0, 0}, 63: []int{0, 0}, 64: []int{0, 0}, 65: []int{0, 0}, 66: []int{0, 0}, 67: []int{0, 0}, 68: []int{0, 0}, 69: []int{140000, 4750000}, 70: []int{4540000, 4110000}, 71: []int{0, 0}, 72: []int{0, 0}, 73: []int{0, 0}, 74: []int{0, 0}, 75: []int{0, 0}, 76: []int{0, 0}, 77: []int{0, 0}, 78: []int{0, 0}, 79: []int{0, 0}, 80: []int{0, 0}, 81: []int{0, 0}, 82: []int{0, 0}, 83: []int{0, 0}, 84: []int{0, 0}, 85: []int{0, 0}, 86: []int{0, 0}, 87: []int{0, 0}, 88: []int{0, 0}, 89: []int{0, 0}, 90: []int{0, 0}, 91: []int{4380000, 1270000}, 92: []int{0, 0}, 93: []int{0, 0}, 94: []int{0, 0}, 95: []int{0, 0}, 96: []int{2280000, 1570000}, 97: []int{60000, 1800000}, 98: []int{3950000, 310000}, 99: []int{440000, 2200000}, 100: []int{0, 0}, 101: []int{0, 0}, 102: []int{0, 0}, 103: []int{0, 0}, 104: []int{0, 0}, 105: []int{0, 0}, 106: []int{0, 0}, 107: []int{0, 0}, 108: []int{0, 0}, 109: []int{0, 0}, 110: []int{0, 0}, 111: []int{0, 0}, 112: []int{0, 0}, 113: []int{0, 0}, 114: []int{0, 0}, 115: []int{0, 0}, 116: []int{0, 0}, 117: []int{0, 0}, 118: []int{0, 0}, 119: []int{0, 0}, 120: []int{0, 0}, 121: []int{0, 0}, 122: []int{0, 0}, 123: []int{0, 0}, 124: []int{0, 0}, 125: []int{0, 0}, 126: []int{0, 0}, 127: []int{0, 0}, 128: []int{0, 0}, 129: []int{0, 0}, 130: []int{0, 0}, 131: []int{0, 0}, 132: []int{0, 0}, 133: []int{0, 0}, 134: []int{0, 0}, 135: []int{0, 0}, 136: []int{0, 0}, 137: []int{0, 0}, 138: []int{0, 0}, 139: []int{0, 0}, 140: []int{0, 0}, 141: []int{0, 0}, 142: []int{0, 0}, 143: []int{0, 0}, 144: []int{0, 0}, 145: []int{0, 0}, 146: []int{0, 0}, 147: []int{0, 0}, 148: []int{0, 0}, 149: []int{0, 0}, 150: []int{0, 0}, 151: []int{0, 0}, 152: []int{0, 0}, 153: []int{0, 0}, 154: []int{0, 0}, 155: []int{0, 0}, 156: []int{0, 0}, 157: []int{0, 0}, 158: []int{0, 0}, 159: []int{0, 0}, 160: []int{0, 0}, 161: []int{0, 0}, 162: []int{0, 0}, 163: []int{0, 0}, 164: []int{0, 0}, 165: []int{0, 0}, 166: []int{0, 0}, 167: []int{1140000, 4210000}, 168: []int{0, 0}, 169: []int{0, 0}, 170: []int{0, 0}, 171: []int{0, 0}, 172: []int{0, 0}, 173: []int{0, 0}, 174: []int{0, 0}, 175: []int{0, 0}, 176: []int{0, 0}, 177: []int{0, 0}, 178: []int{0, 0}, 179: []int{0, 0}, 180: []int{0, 0}, 181: []int{0, 0}, 182: []int{0, 0}, 183: []int{0, 0}, 184: []int{0, 0}, 185: []int{0, 0}, 186: []int{0, 0}, 187: []int{0, 0}, 188: []int{0, 0}, 189: []int{0, 0}, 190: []int{0, 0}, 191: []int{0, 0}, 192: []int{0, 0}, 193: []int{0, 0}, 194: []int{950000, 800000}, 195: []int{2940000, 4440000}, 196: []int{4800000, 470000}, 197: []int{0, 0}, 198: []int{4670000, 3980000}, 199: []int{0, 0}, 200: []int{0, 0}, 201: []int{0, 0}, 202: []int{0, 0}, 203: []int{0, 0}, 204: []int{0, 0}, 205: []int{2550000, 1870000}, 206: []int{1700000, 4630000}, 207: []int{0, 0}, 208: []int{0, 0}, 209: []int{0, 0}, 210: []int{0, 0}, 211: []int{0, 0}, 212: []int{0, 0}, 213: []int{4650000, 3640000}, 214: []int{1520000, 880000}, 215: []int{570000, 3230000}, 216: []int{370000, 4600000}, 217: []int{0, 0}, 218: []int{0, 0}, 219: []int{0, 0}, 220: []int{0, 0}, 221: []int{0, 0}, 222: []int{0, 0}, 223: []int{0, 0}, 224: []int{0, 0}, 225: []int{0, 0}, 226: []int{0, 0}, 227: []int{0, 0}, 228: []int{0, 0}, 229: []int{0, 0}, 230: []int{0, 0}, 231: []int{0, 0}, 232: []int{0, 0}, 233: []int{0, 0}, 234: []int{0, 0}, 235: []int{0, 0}, 236: []int{0, 0}, 237: []int{0, 0}, 238: []int{0, 0}, 239: []int{0, 0}, 240: []int{0, 0}, 241: []int{0, 0}, 242: []int{0, 0}, 243: []int{0, 0}, 244: []int{3880000, 4540000}, 245: []int{2480000, 720000}, 246: []int{0, 0}, 247: []int{0, 0}, 248: []int{0, 0}, 249: []int{0, 0}, 250: []int{0, 0}, 251: []int{0, 0}, 252: []int{0, 0}, 253: []int{0, 0}, 254: []int{0, 0}, 255: []int{0, 0}, 256: []int{0, 0}, 257: []int{0, 0}, 258: []int{0, 0}, 259: []int{0, 0}, 260: []int{0, 0}, 261: []int{0, 0}, 262: []int{0, 0}, 263: []int{0, 0}, 264: []int{0, 0}, 265: []int{0, 0}, 266: []int{0, 0}, 267: []int{0, 0}, 268: []int{2650000, 600000}, 269: []int{1460000, 2160000}, 270: []int{710000, 4970000}, 271: []int{0, 0}, 272: []int{0, 0}, 273: []int{0, 0}, 274: []int{0, 0}, 275: []int{0, 0}, 276: []int{0, 0}, 277: []int{0, 0}, 278: []int{0, 0}, 279: []int{0, 0}, 280: []int{0, 0}, 281: []int{0, 0}, 282: []int{0, 0}, 283: []int{0, 0}, 284: []int{0, 0}, 285: []int{0, 0}, 286: []int{0, 0}, 287: []int{0, 0}, 288: []int{0, 0}, 289: []int{0, 0}, 290: []int{0, 0}, 291: []int{0, 0}, 292: []int{0, 0}, 293: []int{0, 0}, 294: []int{0, 0}, 295: []int{0, 0}, 296: []int{0, 0}, 297: []int{0, 0}, 298: []int{0, 0}, 299: []int{0, 0}, 300: []int{0, 0}, 301: []int{0, 0}, 302: []int{0, 0}, 303: []int{0, 0}, 304: []int{0, 0}, 305: []int{0, 0}, 306: []int{0, 0}, 307: []int{0, 0}, 308: []int{0, 0}, 309: []int{0, 0}, 310: []int{0, 0}, 311: []int{0, 0}, 312: []int{0, 0}, 313: []int{0, 0}, 314: []int{0, 0}, 315: []int{0, 0}, 316: []int{0, 0}, 317: []int{0, 0}, 318: []int{0, 0}, 319: []int{0, 0}, 320: []int{0, 0}, 321: []int{2750000, 3850000}, 322: []int{4110000, 1040000}, 323: []int{0, 0}, 324: []int{0, 0}, 325: []int{0, 0}, 326: []int{0, 0}, 327: []int{0, 0}, 328: []int{0, 0}, 329: []int{3040000, 4990000}, 330: []int{0, 0}, 331: []int{0, 0}, 332: []int{0, 0}, 333: []int{0, 0}, 334: []int{0, 0}, 335: []int{0, 0}, 336: []int{0, 0}, 337: []int{0, 0}, 338: []int{0, 0}, 339: []int{0, 0}, 340: []int{0, 0}, 341: []int{0, 0}, 342: []int{0, 0}, 343: []int{0, 0}, 344: []int{0, 0}, 345: []int{0, 0}, 346: []int{0, 0}, 347: []int{0, 0}, 348: []int{0, 0}, 349: []int{4960000, 4070000}, 350: []int{0, 0}, 351: []int{0, 0}, 352: []int{0, 0}, 353: []int{0, 0}, 354: []int{0, 0}, 355: []int{0, 0}, 356: []int{0, 0}, 357: []int{0, 0}, 358: []int{2060000, 4010000}, 359: []int{0, 0}, 360: []int{3960000, 4550000}, 361: []int{1560000, 650000}, 362: []int{0, 0}, 363: []int{0, 0}, 364: []int{0, 0}, 365: []int{0, 0}, 366: []int{0, 0}, 367: []int{0, 0}, 368: []int{0, 0}, 369: []int{0, 0}, 370: []int{0, 0}, 371: []int{0, 0}, 372: []int{0, 0}, 373: []int{140000, 1510000}, 374: []int{1900000, 1200000}, 375: []int{0, 0}, 376: []int{0, 0}, 377: []int{0, 0}, 378: []int{0, 0}, 379: []int{0, 0}, 380: []int{0, 0}, 381: []int{1780000, 3660000}, 382: []int{3810000, 2000000}, 383: []int{0, 0}, 384: []int{0, 0}, 385: []int{0, 0}, 386: []int{0, 0}, 387: []int{0, 0}, 388: []int{0, 0}, 389: []int{0, 0}, 390: []int{0, 0}, 391: []int{0, 0}, 392: []int{0, 0}, 393: []int{0, 0}, 394: []int{0, 0}, 395: []int{0, 0}, 396: []int{0, 0}, 397: []int{0, 0}, 398: []int{0, 0}, 399: []int{0, 0}, 400: []int{0, 0}, 401: []int{0, 0}, 402: []int{0, 0}, 403: []int{0, 0}, 404: []int{0, 0}, 405: []int{0, 0}, 406: []int{0, 0}, 407: []int{0, 0}, 408: []int{0, 0}, 409: []int{0, 0}, 410: []int{0, 0}, 411: []int{1880000, 2650000}, 412: []int{2800000, 420000}, 413: []int{4420000, 2250000}, 414: []int{4460000, 3020000}, 415: []int{0, 0}, 416: []int{0, 0}, 417: []int{0, 0}, 418: []int{0, 0}, 419: []int{0, 0}, 420: []int{0, 0}, 421: []int{0, 0}, 422: []int{0, 0}, 423: []int{0, 0}, 424: []int{0, 0}, 425: []int{0, 0}, 426: []int{0, 0}, 427: []int{0, 0}, 428: []int{0, 0}, 429: []int{0, 0}, 430: []int{0, 0}, 431: []int{0, 0}, 432: []int{0, 0}, 433: []int{0, 0}, 434: []int{0, 0}, 435: []int{0, 0}, 436: []int{0, 0}, 437: []int{0, 0}, 438: []int{0, 0}, 439: []int{0, 0}, 440: []int{0, 0}, 441: []int{0, 0}, 442: []int{0, 0}, 443: []int{0, 0}, 444: []int{0, 0}, 445: []int{0, 0}, 446: []int{0, 0}, 447: []int{0, 0}, 448: []int{0, 0}, 449: []int{0, 0}, 450: []int{0, 0}, 451: []int{0, 0}, 452: []int{0, 0}, 453: []int{0, 0}, 454: []int{0, 0}, 455: []int{0, 0}, 456: []int{1180000, 4730000}, 457: []int{1850000, 4760000}, 458: []int{0, 0}, 459: []int{0, 0}, 460: []int{4680000, 3160000}, 461: []int{560000, 4420000}, 462: []int{2720000, 2110000}, 463: []int{4010000, 1270000}, 464: []int{0, 0}, 465: []int{970000, 2020000}, 466: []int{580000, 2320000}, 467: []int{850000, 3360000}, 468: []int{3470000, 3290000}, 469: []int{0, 0}, 470: []int{0, 0}, 471: []int{0, 0}, 472: []int{0, 0}, 473: []int{0, 0}, 474: []int{0, 0}, 475: []int{0, 0}, 476: []int{0, 0}, 477: []int{0, 0}, 478: []int{0, 0}, 479: []int{0, 0}, 480: []int{0, 0}, 481: []int{0, 0}, 482: []int{0, 0}, 483: []int{0, 0}, 484: []int{0, 0}, 485: []int{0, 0}, 486: []int{0, 0}, 487: []int{0, 0}, 488: []int{0, 0}, 489: []int{0, 0}, 490: []int{0, 0}, 491: []int{0, 0}, 492: []int{0, 0}, 493: []int{0, 0}, 494: []int{0, 0}, 495: []int{0, 0}, 496: []int{0, 0}, 497: []int{0, 0}, 498: []int{0, 0}, 499: []int{0, 0}, 500: []int{0, 0}, 501: []int{0, 0}, 502: []int{0, 0}, 503: []int{0, 0}, 504: []int{0, 0}, 505: []int{0, 0}, 506: []int{0, 0}, 507: []int{0, 0}, 508: []int{0, 0}, 509: []int{0, 0}, 510: []int{0, 0}, 511: []int{0, 0}, 512: []int{0, 0}, 513: []int{0, 0}, 514: []int{0, 0}, 515: []int{0, 0}, 516: []int{0, 0}, 517: []int{0, 0}, 518: []int{0, 0}, 519: []int{0, 0}, 520: []int{0, 0}, 521: []int{0, 0}, 522: []int{0, 0}, 523: []int{0, 0}, 524: []int{0, 0}, 525: []int{220000, 120000}, 526: []int{0, 0}, 527: []int{0, 0}, 528: []int{0, 0}, 529: []int{0, 0}, 530: []int{0, 0}, 531: []int{0, 0}, 532: []int{0, 0}, 533: []int{0, 0}, 534: []int{0, 0}, 535: []int{0, 0}, 536: []int{0, 0}, 537: []int{0, 0}, 538: []int{0, 0}, 539: []int{0, 0}, 540: []int{0, 0}, 541: []int{0, 0}, 542: []int{0, 0}, 543: []int{0, 0}, 544: []int{0, 0}, 545: []int{0, 0}, 546: []int{0, 0}, 547: []int{0, 0}, 548: []int{0, 0}, 549: []int{0, 0}, 550: []int{0, 0}, 551: []int{0, 0}, 552: []int{0, 0}, 553: []int{0, 0}, 554: []int{0, 0}, 555: []int{0, 0}, 556: []int{4790000, 2990000}, 557: []int{2170000, 4070000}, 558: []int{4070000, 4140000}, 559: []int{2580000, 1940000}, 560: []int{820000, 4650000}, 561: []int{4000000, 2750000}, 562: []int{0, 0}, 563: []int{0, 0}, 564: []int{0, 0}, 565: []int{0, 0}, 566: []int{0, 0}, 567: []int{0, 0}, 568: []int{0, 0}, 569: []int{0, 0}, 570: []int{0, 0}, 571: []int{0, 0}, 572: []int{0, 0}, 573: []int{4650000, 3120000}, 574: []int{2120000, 670000}, 575: []int{2370000, 3190000}, 576: []int{2300000, 1090000}, 577: []int{0, 0}, 578: []int{0, 0}, 579: []int{0, 0}, 580: []int{0, 0}, 581: []int{0, 0}, 582: []int{0, 0}, 583: []int{0, 0}, 584: []int{0, 0}, 585: []int{0, 0}, 586: []int{0, 0}, 587: []int{0, 0}, 588: []int{0, 0}, 589: []int{0, 0}, 590: []int{0, 0}, 591: []int{0, 0}, 592: []int{0, 0}, 593: []int{0, 0}, 594: []int{0, 0}, 595: []int{0, 0}, 596: []int{0, 0}, 597: []int{0, 0}, 598: []int{0, 0}, 599: []int{0, 0}, 600: []int{3930000, 2860000}, 601: []int{2470000, 1940000}, 602: []int{3580000, 3370000}, 603: []int{2450000, 4310000}, 604: []int{0, 0}, 605: []int{0, 0}, 606: []int{0, 0}, 607: []int{0, 0}, 608: []int{0, 0}, 609: []int{0, 0}, 610: []int{0, 0}, 611: []int{0, 0}, 612: []int{0, 0}, 613: []int{0, 0}, 614: []int{0, 0}, 615: []int{0, 0}, 616: []int{0, 0}, 617: []int{0, 0}, 618: []int{0, 0}, 619: []int{0, 0}, 620: []int{0, 0}, 621: []int{0, 0}, 622: []int{0, 0}, 623: []int{0, 0}, 624: []int{0, 0}, 625: []int{4140000, 4390000}, 626: []int{0, 0}, 627: []int{0, 0}, 628: []int{0, 0}, 629: []int{0, 0}, 630: []int{0, 0}, 631: []int{0, 0}, 632: []int{0, 0}, 633: []int{0, 0}, 634: []int{0, 0}, 635: []int{0, 0}, 636: []int{2970000, 3620000}, 637: []int{0, 0}, 638: []int{0, 0}, 639: []int{0, 0}, 640: []int{0, 0}, 641: []int{0, 0}, 642: []int{0, 0}, 643: []int{0, 0}, 644: []int{0, 0}, 645: []int{0, 0}, 646: []int{0, 0}, 647: []int{0, 0}, 648: []int{0, 0}, 649: []int{0, 0}, 650: []int{0, 0}, 651: []int{0, 0}, 652: []int{10000, 2350000}, 653: []int{3670000, 2810000}, 654: []int{240000, 4190000}, 655: []int{0, 0}, 656: []int{0, 0}, 657: []int{0, 0}, 658: []int{0, 0}, 659: []int{3560000, 2110000}, 660: []int{860000, 1590000}, 661: []int{4170000, 2220000}, 662: []int{0, 0}, 663: []int{0, 0}, 664: []int{0, 0}, 665: []int{0, 0}, 666: []int{0, 0}, 667: []int{4910000, 4020000}, 668: []int{1280000, 1900000}, 669: []int{90000, 3660000}, 670: []int{0, 0}, 671: []int{0, 0}, 672: []int{0, 0}, 673: []int{0, 0}, 674: []int{0, 0}, 675: []int{0, 0}, 676: []int{0, 0}, 677: []int{0, 0}, 678: []int{0, 0}, 679: []int{0, 0}, 680: []int{0, 0}, 681: []int{0, 0}, 682: []int{0, 0}, 683: []int{0, 0}, 684: []int{0, 0}, 685: []int{0, 0}, 686: []int{0, 0}, 687: []int{0, 0}, 688: []int{4540000, 4460000}, 689: []int{2900000, 1170000}, 690: []int{1920000, 3320000}, 691: []int{110000, 2740000}, 692: []int{0, 0}, 693: []int{0, 0}, 694: []int{0, 0}, 695: []int{0, 0}, 696: []int{0, 0}, 697: []int{0, 0}, 698: []int{0, 0}, 699: []int{0, 0}, 700: []int{0, 0}, 701: []int{0, 0}, 702: []int{0, 0}, 703: []int{0, 0}, 704: []int{0, 0}, 705: []int{0, 0}, 706: []int{0, 0}, 707: []int{0, 0}, 708: []int{0, 0}, 709: []int{0, 0}, 710: []int{0, 0}, 711: []int{0, 0}, 712: []int{0, 0}, 713: []int{0, 0}, 714: []int{0, 0}, 715: []int{0, 0}, 716: []int{0, 0}, 717: []int{0, 0}, 718: []int{3470000, 3350000}, 719: []int{2510000, 3220000}, 720: []int{360000, 880000}, 721: []int{940000, 1360000}, 722: []int{0, 0}, 723: []int{0, 0}, 724: []int{0, 0}, 725: []int{0, 0}, 726: []int{0, 0}, 727: []int{0, 0}, 728: []int{0, 0}, 729: []int{0, 0}, 730: []int{0, 0}, 731: []int{0, 0}, 732: []int{0, 0}, 733: []int{0, 0}, 734: []int{0, 0}, 735: []int{0, 0}, 736: []int{3710000, 4570000}, 737: []int{1760000, 1830000}, 738: []int{1920000, 4860000}, 739: []int{2560000, 1580000}, 740: []int{0, 0}, 741: []int{0, 0}, 742: []int{0, 0}, 743: []int{0, 0}, 744: []int{0, 0}, 745: []int{0, 0}, 746: []int{0, 0}, 747: []int{0, 0}, 748: []int{4350000, 200000}, 749: []int{1140000, 2250000}, 750: []int{4570000, 100000}, 751: []int{0, 0}, 752: []int{0, 0}, 753: []int{0, 0}, 754: []int{0, 0}, 755: []int{90000, 880000}, 756: []int{110000, 4320000}, 757: []int{0, 0}, 758: []int{0, 0}, 759: []int{0, 0}, 760: []int{4270000, 4680000}, 761: []int{1700000, 1100000}, 762: []int{2630000, 2860000}, 763: []int{0, 0}, 764: []int{0, 0}, 765: []int{0, 0}, 766: []int{0, 0}, 767: []int{0, 0}, 768: []int{0, 0}, 769: []int{0, 0}, 770: []int{0, 0}, 771: []int{0, 0}, 772: []int{0, 0}, 773: []int{0, 0}, 774: []int{0, 0}, 775: []int{0, 0}, 776: []int{0, 0}, 777: []int{1400000, 2880000}, 778: []int{0, 0}, 779: []int{0, 0}, 780: []int{0, 0}, 781: []int{0, 0}, 782: []int{4230000, 3790000}, 783: []int{0, 0}, 784: []int{0, 0}, 785: []int{0, 0}, 786: []int{0, 0}, 787: []int{0, 0}, 788: []int{0, 0}, 789: []int{0, 0}, 790: []int{0, 0}, 791: []int{0, 0}, 792: []int{0, 0}, 793: []int{4680000, 3240000}, 794: []int{0, 0}, 795: []int{0, 0}, 796: []int{0, 0}, 797: []int{0, 0}, 798: []int{0, 0}, 799: []int{0, 0}, 800: []int{2570000, 520000}, 801: []int{3710000, 2760000}, 802: []int{4920000, 4850000}, 803: []int{4210000, 2980000}, 804: []int{0, 0}, 805: []int{1860000, 680000}, 806: []int{840000, 2670000}, 807: []int{4490000, 3560000}, 808: []int{0, 0}, 809: []int{0, 0}, 810: []int{0, 0}, 811: []int{0, 0}, 812: []int{0, 0}, 813: []int{0, 0}, 814: []int{0, 0}, 815: []int{0, 0}, 816: []int{0, 0}, 817: []int{0, 0}, 818: []int{0, 0}, 819: []int{0, 0}, 820: []int{0, 0}, 821: []int{0, 0}, 822: []int{0, 0}, 823: []int{0, 0}, 824: []int{0, 0}, 825: []int{0, 0}, 826: []int{0, 0}, 827: []int{0, 0}, 828: []int{0, 0}, 829: []int{0, 0}, 830: []int{0, 0}, 831: []int{0, 0}, 832: []int{0, 0}, 833: []int{0, 0}, 834: []int{0, 0}, 835: []int{0, 0}, 836: []int{0, 0}, 837: []int{0, 0}, 838: []int{0, 0}, 839: []int{0, 0}, 840: []int{0, 0}, 841: []int{0, 0}, 842: []int{0, 0}, 843: []int{0, 0}, 844: []int{0, 0}, 845: []int{0, 0}, 846: []int{0, 0}, 847: []int{0, 0}, 848: []int{0, 0}, 849: []int{0, 0}, 850: []int{0, 0}, 851: []int{0, 0}, 852: []int{900000, 3960000}, 853: []int{2520000, 4220000}, 854: []int{0, 0}, 855: []int{0, 0}, 856: []int{0, 0}, 857: []int{0, 0}, 858: []int{0, 0}, 859: []int{0, 0}, 860: []int{0, 0}, 861: []int{0, 0}, 862: []int{0, 0}, 863: []int{0, 0}, 864: []int{0, 0}, 865: []int{0, 0}, 866: []int{0, 0}, 867: []int{0, 0}, 868: []int{0, 0}, 869: []int{230000, 4040000}, 870: []int{1980000, 900000}, 871: []int{1760000, 3560000}, 872: []int{0, 0}, 873: []int{0, 0}, 874: []int{0, 0}, 875: []int{4360000, 220000}, 876: []int{280000, 3960000}, 877: []int{0, 0}, 878: []int{0, 0}, 879: []int{0, 0}, 880: []int{0, 0}, 881: []int{0, 0}, 882: []int{0, 0}, 883: []int{0, 0}, 884: []int{0, 0}, 885: []int{0, 0}, 886: []int{0, 0}, 887: []int{0, 0}, 888: []int{0, 0}, 889: []int{0, 0}, 890: []int{0, 0}, 891: []int{0, 0}, 892: []int{0, 0}, 893: []int{0, 0}, 894: []int{0, 0}, 895: []int{0, 0}, 896: []int{0, 0}, 897: []int{1630000, 930000}, 898: []int{0, 0}, 899: []int{0, 0}, 900: []int{0, 0}, 901: []int{0, 0}, 902: []int{0, 0}, 903: []int{0, 0}, 904: []int{0, 0}, 905: []int{0, 0}, 906: []int{0, 0}, 907: []int{0, 0}, 908: []int{2110000, 2570000}, 909: []int{0, 0}, 910: []int{0, 0}, 911: []int{0, 0}, 912: []int{0, 0}, 913: []int{0, 0}, 914: []int{0, 0}, 915: []int{0, 0}, 916: []int{4410000, 4060000}, 917: []int{0, 0}, 918: []int{0, 0}, 919: []int{0, 0}, 920: []int{0, 0}, 921: []int{0, 0}, 922: []int{0, 0}, 923: []int{0, 0}, 924: []int{0, 0}, 925: []int{0, 0}, 926: []int{0, 0}, 927: []int{0, 0}, 928: []int{0, 0}, 929: []int{0, 0}, 930: []int{0, 0}, 931: []int{0, 0}, 932: []int{0, 0}, 933: []int{0, 0}, 934: []int{0, 0}, 935: []int{0, 0}, 936: []int{0, 0}, 937: []int{0, 0}, 938: []int{0, 0}, 939: []int{0, 0}, 940: []int{0, 0}, 941: []int{0, 0}, 942: []int{0, 0}, 943: []int{0, 0}, 944: []int{0, 0}, 945: []int{0, 0}, 946: []int{2150000, 1580000}, 947: []int{3520000, 3500000}, 948: []int{0, 0}, 949: []int{0, 0}, 950: []int{0, 0}, 951: []int{0, 0}, 952: []int{0, 0}, 953: []int{0, 0}, 954: []int{0, 0}, 955: []int{0, 0}, 956: []int{0, 0}, 957: []int{0, 0}, 958: []int{0, 0}, 959: []int{0, 0}, 960: []int{0, 0}, 961: []int{0, 0}, 962: []int{0, 0}, 963: []int{0, 0}, 964: []int{0, 0}, 965: []int{0, 0}, 966: []int{0, 0}, 967: []int{0, 0}, 968: []int{0, 0}, 969: []int{0, 0}, 970: []int{0, 0}, 971: []int{2330000, 2810000}, 972: []int{0, 0}, 973: []int{0, 0}, 974: []int{0, 0}, 975: []int{0, 0}, 976: []int{0, 0}, 977: []int{0, 0}, 978: []int{0, 0}, 979: []int{0, 0}, 980: []int{0, 0}, 981: []int{0, 0}, 982: []int{0, 0}, 983: []int{0, 0}, 984: []int{0, 0}, 985: []int{0, 0}, 986: []int{0, 0}, 987: []int{0, 0}, 988: []int{0, 0}, 989: []int{0, 0}, 990: []int{0, 0}, 991: []int{0, 0}, 992: []int{0, 0}, 993: []int{0, 0}, 994: []int{0, 0}, 995: []int{0, 0}, 996: []int{0, 0}, 997: []int{0, 0}, 998: []int{0, 0}, 999: []int{0, 0}, 1000: []int{0, 0}},
		}
		fmt.Println("dasda")
		dynamicClassing := treeMerge(sizeEnd, 25, t)
		totalCost := 0.0
		for _, elem := range dynamicClassing {
			//fmt.Println("Class name: ", elem.className, " Total days: ", elem.days, " Cost: ", elem.cost)
			totalCost += elem.cost
		}
		fmt.Println("Total Cost for Dynamic Classes: ", totalCost)

		tiering := calculateTiering(t)
		tieringCost := 0.0
		for _, elem := range tiering {
			//fmt.Println("Class name: ", elem.className, " Total days: ", elem.days, " Cost: ", elem.cost)
			tieringCost += elem.cost
		}
		fmt.Println("Total Cost for Intelligent Tiering: ", tieringCost)

	//	retorno := optimizer(1, sizeEnd)
	//	minimo := math.MaxFloat64
	//	var minSetup []ObjClass
	//	for _, interval := range retorno {
	//		var cal float64
	//		var first int
	//		var last int
	//		var dynamicClassing []ObjClass
	//		for _, subInt := range interval {
	//			if len(subInt) == 1 {
	//				first = subInt[0]
	//				last = subInt[0]
	//			} else {
	//				first = subInt[0]
	//				last = subInt[1]
	//			}
	//			dynamicClassing = append(dynamicClassing, treeMerge(first, last, traces.TraceArray[0])...)
	//			for _, elem := range dynamicClassing {
	//				cal += elem.cost
	//			}
	//
	//		}
	//		if cal < minimo {
	//			minimo = cal
	//			minSetup = dynamicClassing
	//		}
	//
	//	}
	//	fmt.Println(minSetup, minimo)
	//
		fmt.Fprintf(w, "%d\n", tieringCost, totalCost)
	}
}
//func optimizer(windowStart int, windowEnd int)  [][][]int {
//	return createIntervals0(windowStart, 1000, nil)
//
//}
//func createIntervals0(windowStart int, windowEnd int, intervals [][][]int)  [][][]int {
//	intervals = createList(windowStart, windowEnd, intervals)
//	var result [][][]int
//	for _, k := range intervals{
//		//fmt.Println(result)
//
//		interval:=createIntervals(windowStart, windowEnd, [][][]int{k})
//		fmt.Println("========================================================")
//		result = append(result, interval...)
//		//fmt.Println("comeco",k, createIntervals(windowStart, windowEnd, [][][]int{k}))
//
//	}
//	return result
//
//}
//
//func createIntervals(windowStart int, windowEnd int, intervals [][][]int) [][][]int {
//	if intervals == nil{
//		intervals = createList(windowStart, windowEnd, intervals)
//	}
//	//fmt.Println("frist",intervals)
//	i := 0
//
//	//control := true
//	for true {
//		var newSubIntervals [][][]int
//		i +=1
//		if  len(intervals) == 0{
//			return nil
//		}
//		for i, elem := range intervals {
//				intervals = RemoveIndex(intervals, i)
//			var temp [][][]int
//			var lastElem int
//			if len(elem[len(elem)-1]) == 1 {
//				lastElem = elem[len(elem)-1][0]+1
//			} else {
//				lastElem = elem[len(elem)-1][1]+1
//			}
//			//fmt.Println("lastElem", lastElem)
//			if lastElem <= windowEnd {
//				temp = createList(lastElem, windowEnd, intervals)
//			}
//
//			if temp != nil && len(temp) > 0{
//				//fmt.Println("temp", temp)
//				for _, newElem := range temp {
//					//fmt.Println("loopbefore", elem, newElem[0])
//					//fmt.Println("intervals", newSubIntervals)
//					temp:= [][][]int{append(elem, newElem[0])}
//					retorno := createIntervals(windowStart, windowEnd, temp)
//					//fmt.Println("intervals after", newSubIntervals, retorno)
//					//fmt.Println("intervals", append(newSubIntervals, retorno...))
//					newSubIntervals = append(newSubIntervals, retorno...)
//					//fmt.Println("loopafeter", newSubIntervals)
//				}
//				//fmt.Println("aqui")
//
//				intervals = append(intervals, newSubIntervals...)
//				//fmt.Println(intervals)
//				return intervals
//			} else {
//				//fmt.Println("second return ", [][][]int{elem})
//				return [][][]int{elem}
//
//			}
//		}
//		//fmt.Println("end")
//	}
//	//fmt.Println("end2")
//	//fmt.Println(intervals)
//	return intervals
//}
//
//func createList(windowStart int, windowEnd int, intervals [][][]int) [][][]int {
//	if windowStart > windowEnd {
//		return nil
//	}
//	intervals = [][][]int{}
//	sliceOfSlices := make([][]int, 1)
//	sliceOfSlices[0] = []int{windowStart}
//	intervals = append(intervals, sliceOfSlices)
//	for i := windowStart + 1; i <= windowEnd; i++ {
//		sliceOfSlices := make([][]int, 1)
//		sliceOfSlices[0] = []int{windowStart, i}
//		intervals = append(intervals, sliceOfSlices)
//	}
//	return intervals
//
//}

func calculateTiering (traceArray ObjTrace) []ObjClass {
	daysCounter := 0
	costObjects := []ObjClass{{ className: "standard", size: traceArray.ObjSize}}
	currentObj := 0
	for i:= 0; i < len(traceArray.Requests); i++ {
		req := traceArray.Requests[i+1]
		costObjects[currentObj].days += 1
		costObjects[currentObj].BillingDays += 1
		costObjects[currentObj].daysCountingOntBill += 1
		if req[0] != 0 || req[1] != 0 {

			costObjects[currentObj].getRequests += float64(req[0])
			costObjects[currentObj].postRequests += float64(req[1])

			if costObjects[currentObj].className != "standard"{
				costObjects = append(costObjects, ObjClass{className: "standard", size: traceArray.ObjSize, days: 1})
				currentObj += 1
			}
			daysCounter = 0
		} else {
			daysCounter += 1
			if daysCounter >= 30 && costObjects[currentObj].className != "archive" {

				newClass := ""
				if costObjects[currentObj].className == "standard" {
					newClass = "infrequent access"
				} else if costObjects[currentObj].className == "infrequent access" { newClass = "glacier"
				} else if costObjects[currentObj].className == "glacier" { newClass = "archive" }
				currentObj += 1
				daysCounter = 0
				costObjects = append(costObjects, ObjClass{className: newClass, size: traceArray.ObjSize, retrievalRequests: []float64{0, 0, 0}})
			}
		}
	}

	for i, obj := range costObjects{
		var previousObj ObjClass
		if i > 0 {
			previousObj = costObjects[i-1]

		}
		costObjects[i].cost += calculateObjCostForClass(getClass(obj.className), obj, previousObj)
	}
	return costObjects
}

func treeMerge(lastWindowDay int, windowSize int, traceArray ObjTrace) []ObjClass{
	currentDay := 1
	currentObject := calculate(currentDay, currentDay + windowSize, traceArray, ObjClass{})
	currentDay += windowSize
	costObjects := []ObjClass{currentObject}
	for true {
		if currentDay <= lastWindowDay {

			currentObject = calculate(currentDay, int(math.Min(float64(currentDay+windowSize), float64(lastWindowDay))), traceArray, currentObject)
			costObjects = append(costObjects, currentObject)
		} else { break }
		currentDay += windowSize
	}
	return costObjects
	//middleWindowDay := (float64(lastWindowDay) + float64(firstWindowDay)) / 2.0
	//if lastWindowDay - firstWindowDay >= windowSize {
	//
	//	treeMerge(firstWindowDay, int(math.Floor(middleWindowDay)), windowSize, traceArray)
	//
	//	treeMerge(int(math.Ceil(middleWindowDay)),  lastWindowDay, windowSize, traceArray)
	//
	//	return minCost(treeMerge(firstWindowDay, int(math.Floor(middleWindowDay)), windowSize, traceArray),treeMerge(int(math.Ceil(middleWindowDay)), lastWindowDay, windowSize, traceArray), traceArray)
	//} else {
	//
	//	return []ObjClass{calculate(firstWindowDay, lastWindowDay, traceArray)}
	//}
}
func getClass(className string) S3Class {
	for _, class := range classes {
		if class.name == className {
			return class
		}
	}
	return S3Class{}
}


func calculate(firstDay int, lastDay int, trace ObjTrace, previousClass ObjClass) ObjClass {

	postInInterval, getInInterval, retrievesInInterval := sumRequests(firstDay, lastDay, trace)

	prices := []float64{-0.1, -0.1, -0.1, -0.1}
	finalClass := ObjClass{size: trace.ObjSize, getRequests: getInInterval, postRequests: postInInterval, retrievalRequests: retrievesInInterval, days: lastDay - firstDay, daysCountingOntBill: lastDay - firstDay}
	for i, class := range classes {
		prices[i] = calculateObjCostForClass(class, finalClass, previousClass) // fazer ajutes do data retrieve para infrequent accesd
	}

	className, minPrice := FindMin(prices)
	finalClass.className = className
	finalClass.cost = minPrice

	if className == previousClass.className {
		finalClass.BillingDays = previousClass.BillingDays + finalClass.days
		classDays := getClass(className).days
		if previousClass.BillingDays <  classDays{
			finalClass.daysCountingOntBill = int(math.Max(0, float64(previousClass.BillingDays + finalClass.days-classDays)))
		}
	} else {
		finalClass.BillingDays = finalClass.days
	}
	return finalClass
}

func sumRequests(firstDay int, lastDay int, trace ObjTrace) (float64, float64, []float64) {
	postInInterval, getInInterval, retrievesInInterval := 0.0, 0.0, []float64{0.0, 0.0, 0.0}

	for day := firstDay; day <= lastDay; day++ {
		totalRequests, ok := trace.Requests[day]
		if ok {
			getInInterval += float64(totalRequests[0])
			postInInterval += float64(totalRequests[1])
			if totalRequests[0] + totalRequests[1] > 0 {
				if day == 1 {
					retrievesInInterval[0] += 1.0 // expedite
				} else if day == 2 {
					retrievesInInterval[1] += 1.0 // standard
				} else {
					retrievesInInterval[2] += 1.0 // bulk
				}
			}

		}
	}
	return postInInterval, getInInterval, retrievesInInterval
}

func calculateObjCostForClass(class S3Class, object ObjClass, previousObject ObjClass) float64 {
	storage := 0.0
	if previousObject.className != class.name || previousObject.BillingDays >= getClass(previousObject.className).days{
		storage = storagePriceCalculate(class, object)
	}

	retrievalPrice := 0.0 // for glacier classes only
	// adjust for ia class
	if class.name == "glacier" {
		retrievalPrice = object.retrievalRequests[0] * class.retrievalPrice[0] + object.retrievalRequests[1]*class.retrievalPrice[1] + object.retrievalRequests[2]*class.retrievalPrice[2]

	} else if class.name == "archive" { // there's no expedite option for s3 glacier deep archive
		if object.retrievalRequests[0] == 0 {
			retrievalPrice = object.retrievalRequests[1]*class.retrievalPrice[1] + object.retrievalRequests[2]*class.retrievalPrice[2]
		} else {
			retrievalPrice = math.MaxFloat64
		}
	} else if class.name == "infrequent access" {
		retrievalPrice = (object.retrievalRequests[0] + object.retrievalRequests[1] + object.retrievalRequests[2])*class.retrievalData[0] * float64(object.size)
	}

	if (class.name == "glacier" || class.name == "archive") && retrievalPrice != 0 {
		for i, req := range object.retrievalRequests{
			if req != 0 {
				//fmt.Println(float64(object.size) , class.retrievalData[i] , object.retrievalRequests[i])
				retrievalPrice += float64(object.size) * class.retrievalData[i] * object.retrievalRequests[i]
			}
		}
	}
	requests := class.postPrice * object.postRequests + class.getPrice * object.getRequests + retrievalPrice
	//fmt.Println(class.name,storage ,requests, retrievalPrice)
	return storage + requests
}

func storagePriceCalculate(class S3Class, object ObjClass) float64 {
	storage := 0.0
	days := math.Ceil(float64(object.days)/float64(30))
	if class.name == "standard" {
		if days < 1{
			days = 1
		}
		firstClass := 50000.0
		secondClass := 450000.0
		thirdClass := 500000.0
		storage = class.storagePrice[0] * math.Min(float64(object.size), firstClass) // price for the first 50TB
		if float64(object.size) > 50000 {
			storage += class.storagePrice[1] * math.Min(float64(object.size) - firstClass, secondClass) // price for the next 450TB
		}
		if float64(object.size) > 500000 {
			storage += class.storagePrice[2] * math.Min(float64(object.size) - thirdClass, thirdClass) // price for over 500TB
		}
		storage *= days
	} else if class.name == "glacier" {
		if days < 3{
			days = 3
		}
		storage = days * class.storagePrice[0] * float64(object.size) // since storagePrice consider 30 day and the minimum billing is 90
	} else if class.name == "archive" {
		if days < 6{
			days = 6
		}
		storage = days * class.storagePrice[0] * float64(object.size) // since storagePrice consider 30 day and the minimum billing is 180
	}	else {
		if days < 1{
			days = 1
		}
		storage = class.storagePrice[0] * float64(object.size) * days
	}
	return storage
}

func FindMin(prices []float64) (name string, min float64) {
	min = prices[0]
	name = "archive"
	for i, price := range prices {
		if price < min {
			min = price
			name = classes[i].name
		}
	}
	return
}
func RemoveIndex(s [][][]int, index int) [][][]int {

	if len(s) == 0  {
		return nil
	}
	if len(s) - 1 <= index {
		return s[:index]
	}
	if index == 0 {
		return s[index+1:]
	}
	return append(s[:index], s[index+1:]...)
}


//func speculateRight(rightCost []ObjClass, right_i int, leftCost []ObjClass, left_i int, rawPrice float64, lowerClass S3Class) ObjClass {
//	unblocked := false
//	newTraceObject := joinTrace(rightCost[right_i], leftCost[left_i])
//	for _, class := range classes {
//		if class.name == lowerClass.name || unblocked {
//			unblocked = true
//			newPrice := calculateObjCostForClass(lowerClass, newTraceObject)
//			if newPrice < rawPrice {
//				rawPrice = newPrice // verificar se o custo de outros objetos dependem dele
//				newTraceObject.cost = rawPrice
//				newTraceObject.className = class.name
//			}
//		}
//	}
//
//	return newTraceObject
//}
//
//func joinTrace(traceObject ObjClass, traceObject2 ObjClass) ObjClass {
//	newRetrievalRequests := []float64{0.0, 0.0, 0.0}
//
//	for i := range newRetrievalRequests{
//		newRetrievalRequests[i] = traceObject.retrievalRequests[i] + traceObject2.retrievalRequests[i]
//	}
//	return ObjClass{
//		className:           "",
//		days:                traceObject.days + traceObject2.days,
//		daysCountingOntBill: traceObject.days + traceObject2.days,
//		cost:                0,
//		size:                traceObject.size,
//		getRequests:         traceObject.getRequests + traceObject2.getRequests,
//		postRequests:        traceObject.postRequests + traceObject2.postRequests,
//		retrievalRequests:   newRetrievalRequests,
//	}
//
//}


//func minCost(leftCost []ObjClass, rightCost []ObjClass, trace ObjTrace)  []ObjClass {
//	right_i, left_i := 0, len(leftCost) -1
//	if leftCost[left_i].className != rightCost[right_i].className {
//		rawPrice := leftCost[left_i].cost + rightCost[right_i].cost // corrigir o calculo com o intervalo do array
//		lowerClass := getClass(leftCost[left_i].className, rightCost[right_i].className)
//		newObject := speculateRight(rightCost, right_i, leftCost, len(leftCost) - 1, rawPrice, lowerClass)
//		if newObject.className != "" {
//			leftCost[left_i] = newObject
//			rightCost = rightCost[1:]
//		}
//	} else {
//		return eliminateStorageRedundancy(leftCost, rightCost)
//	}
//	return append(leftCost, rightCost...)
//}

//func eliminateStorageRedundancy(rightObjects []ObjClass, leftObjects []ObjClass) []ObjClass{
//	summaryObjects := append(leftObjects, rightObjects...)
//	for i := range summaryObjects {
//
//		if i + 1 < len(summaryObjects) && summaryObjects[i+1].daysCountingOntBill > 0 && summaryObjects[i].className == summaryObjects[i+1].className {
//			class := getClass(summaryObjects[i].className)
//			remainingDays := class.days - summaryObjects[i].daysCountingOntBill
//			if remainingDays > 0  { // se há dias pagos extras
//				if remainingDays < summaryObjects[i+1].days {
//					summaryObjects[i+1].daysCountingOntBill = summaryObjects[i+1].days - remainingDays
//				} else {
//					summaryObjects[i+1].daysCountingOntBill = 0
//					summaryObjects[i+1].cost -= storagePriceCalculate(class, summaryObjects[i+1])
//				}
//			}
//		}
//	}
//	return summaryObjects
//	//// caso haja mudanca no costObjects (left) eh preciso atualiza-lo
//	//if leftObjects[len(leftObjects) - 1].className == rightObjects[0].className {
//	//	class := getClass(leftObjects[len(leftObjects)-1].className, leftObjects[len(leftObjects)-1].className)
//	//	remainingDays := class.days - leftObjects[len(leftObjects) - 1].daysCountingOntBill
//	//	if remainingDays > 0  { // se há dias pagos extras
//	//		if remainingDays < rightObjects[0].days {
//	//			rightObjects[0].daysCountingOntBill = rightObjects[0].days - remainingDays
//	//		} else {
//	//			rightObjects[0].daysCountingOntBill = remainingDays - rightObjects[0].days
//	//			rightObjects[0].cost -= storagePriceCalculate(class, rightObjects[0])
//	//		}
//	//	}
//	//	return eliminateStorageRedundancyRight(rightObjects)
//	//} else {
//	//	return rightObjects
//	//}
//}
//
//func eliminateStorageRedundancyRight(rightObjects []ObjClass) []ObjClass{
//	class := getClass(rightObjects[0].className)
//	remainingDays := class.days - rightObjects[0].daysCountingOntBill
//	if remainingDays > 0 && len(rightObjects) > 1 {
//		index := 1
//		for true {
//			if rightObjects[index-1].className == rightObjects[index].className{
//				if remainingDays < rightObjects[index].days {
//					rightObjects[index].daysCountingOntBill = rightObjects[index].days - remainingDays
//					remainingDays = 0
//				} else {
//					rightObjects[index].daysCountingOntBill = remainingDays - rightObjects[index].days
//					rightObjects[index].cost -= storagePriceCalculate(class, rightObjects[index])
//					remainingDays -= rightObjects[index].days
//				}
//			} else { break }
//			if remainingDays == 0 { break }
//		}
//	}
//	return rightObjects
//}

