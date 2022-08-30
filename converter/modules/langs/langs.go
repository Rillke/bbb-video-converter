package langs

type Language struct {
	Name  string
	Local string
	One   string
	Two   string
	TwoT  string
	TwoB  string
	Three string
}

var LanguageList = map[string]Language{"ab": {"Abkhaz", "Аҧсуа", "ab", "abk", "abk", "abk", "abk"},
	"aa": {"Afar", "Afaraf", "aa", "aar", "aar", "aar", "aar"},
	"af": {"Afrikaans", "Afrikaans", "af", "afr", "afr", "afr", "afr"},
	"ak": {"Akan", "Akan", "ak", "aka", "aka", "aka", "aka"},
	"sq": {"Albanian", "Shqip", "sq", "sqi", "sqi", "alb", "sqi"},
	"am": {"Amharic", "አማርኛ", "am", "amh", "amh", "amh", "amh"},
	"ar": {"Arabic", "العربية", "ar", "ara", "ara", "ara", "ara"},
	"an": {"Aragonese", "Aragonés", "an", "arg", "arg", "arg", "arg"},
	"hy": {"Armenian", "Հայերեն", "hy", "hye", "hye", "arm", "hye"},
	"as": {"Assamese", "অসমীয়া", "as", "asm", "asm", "asm", "asm"},
	"av": {"Avaric", "Авар", "av", "ava", "ava", "ava", "ava"},
	"ae": {"Avestan", "avesta", "ae", "ave", "ave", "ave", "ave"},
	"ay": {"Aymara", "Aymar", "ay", "aym", "aym", "aym", "aym"},
	"az": {"Azerbaijani", "Azərbaycanca", "az", "aze", "aze", "aze", "aze"},
	"bm": {"Bambara", "Bamanankan", "bm", "bam", "bam", "bam", "bam"},
	"ba": {"Bashkir", "Башҡортса", "ba", "bak", "bak", "bak", "bak"},
	"eu": {"Basque", "Euskara", "eu", "eus", "eus", "baq", "eus"},
	"be": {"Belarusian", "Беларуская", "be", "bel", "bel", "bel", "bel"},
	"bn": {"Bengali", "বাংলা", "bn", "ben", "ben", "ben", "ben"},
	"bh": {"Bihari", "भोजपुरी", "bh", "bih", "bih", "bih", "bih"},
	"bi": {"Bislama", "Bislama", "bi", "bis", "bis", "bis", "bis"},
	"bs": {"Bosnian", "Bosanski", "bs", "bos", "bos", "bos", "bos"},
	"br": {"Breton", "Brezhoneg", "br", "bre", "bre", "bre", "bre"},
	"bg": {"Bulgarian", "Български", "bg", "bul", "bul", "bul", "bul"},
	"my": {"Burmese", "မြန်မာဘာသာ", "my", "mya", "mya", "bur", "mya"},
	"ca": {"Catalan", "Català", "ca", "cat", "cat", "cat", "cat"},
	"ch": {"Chamorro", "Chamoru", "ch", "cha", "cha", "cha", "cha"},
	"ce": {"Chechen", "Нохчийн", "ce", "che", "che", "che", "che"},
	"ny": {"Chichewa", "Chichewa", "ny", "nya", "nya", "nya", "nya"},
	"zh": {"Chinese", "中文", "zh", "zho", "zho", "chi", "zho"},
	"cv": {"Chuvash", "Чӑвашла", "cv", "chv", "chv", "chv", "chv"},
	"kw": {"Cornish", "Kernewek", "kw", "cor", "cor", "cor", "cor"},
	"co": {"Corsican", "Corsu", "co", "cos", "cos", "cos", "cos"},
	"cr": {"Cree", "ᓀᐦᐃᔭᐍᐏᐣ", "cr", "cre", "cre", "cre", "cre"},
	"hr": {"Croatian", "Hrvatski", "hr", "hrv", "hrv", "hrv", "hrv"},
	"cs": {"Czech", "Čeština", "cs", "ces", "ces", "cze", "ces"},
	"da": {"Danish", "Dansk", "da", "dan", "dan", "dan", "dan"},
	"dv": {"Divehi", "Divehi", "dv", "div", "div", "div", "div"},
	"nl": {"Dutch", "Nederlands", "nl", "nld", "nld", "dut", "nld"},
	"dz": {"Dzongkha", "རྫོང་ཁ", "dz", "dzo", "dzo", "dzo", "dzo"},
	"en": {"English", "English", "en", "eng", "eng", "eng", "eng"},
	"eo": {"Esperanto", "Esperanto", "eo", "epo", "epo", "epo", "epo"},
	"et": {"Estonian", "Eesti", "et", "est", "est", "est", "est"},
	"ee": {"Ewe", "Eʋegbe", "ee", "ewe", "ewe", "ewe", "ewe"},
	"fo": {"Faroese", "Føroyskt", "fo", "fao", "fao", "fao", "fao"},
	"fj": {"Fijian", "Na Vosa Vaka-Viti", "fj", "fij", "fij", "fij", "fij"},
	"fi": {"Finnish", "Suomi", "fi", "fin", "fin", "fin", "fin"},
	"fr": {"French", "Français", "fr", "fra", "fra", "fre", "fra"},
	"ff": {"Fula", "Fulfulde", "ff", "ful", "ful", "ful", "ful"},
	"gl": {"Galician", "Galego", "gl", "glg", "glg", "glg", "glg"},
	"ka": {"Georgian", "ქართული", "ka", "kat", "kat", "geo", "kat"},
	"de": {"German", "Deutsch", "de", "deu", "deu", "ger", "deu"},
	"el": {"Greek", "Ελληνικά", "el", "ell", "ell", "gre", "ell"},
	"gn": {"Guaraní", "Avañe'ẽ", "gn", "grn", "grn", "grn", "grn"},
	"gu": {"Gujarati", "ગુજરાતી", "gu", "guj", "guj", "guj", "guj"},
	"ht": {"Haitian", "Kreyòl Ayisyen", "ht", "hat", "hat", "hat", "hat"},
	"ha": {"Hausa", "هَوُسَ", "ha", "hau", "hau", "hau", "hau"},
	"he": {"Hebrew", "עברית", "he", "heb", "heb", "heb", "heb"},
	"hz": {"Herero", "Otjiherero", "hz", "her", "her", "her", "her"},
	"hi": {"Hindi", "हिन्दी", "hi", "hin", "hin", "hin", "hin"},
	"ho": {"Hiri Motu", "Hiri Motu", "ho", "hmo", "hmo", "hmo", "hmo"},
	"hu": {"Hungarian", "Magyar", "hu", "hun", "hun", "hun", "hun"},
	"ia": {"Interlingua", "Interlingua", "ia", "ina", "ina", "ina", "ina"},
	"id": {"Indonesian", "Bahasa Indonesia", "id", "ind", "ind", "ind", "ind"},
	"ie": {"Interlingue", "Interlingue", "ie", "ile", "ile", "ile", "ile"},
	"ga": {"Irish", "Gaeilge", "ga", "gle", "gle", "gle", "gle"},
	"ig": {"Igbo", "Igbo", "ig", "ibo", "ibo", "ibo", "ibo"},
	"ik": {"Inupiaq", "Iñupiak", "ik", "ipk", "ipk", "ipk", "ipk"},
	"io": {"Ido", "Ido", "io", "ido", "ido", "ido", "ido"},
	"is": {"Icelandic", "Íslenska", "is", "isl", "isl", "ice", "isl"},
	"it": {"Italian", "Italiano", "it", "ita", "ita", "ita", "ita"},
	"iu": {"Inuktitut", "ᐃᓄᒃᑎᑐᑦ", "iu", "iku", "iku", "iku", "iku"},
	"ja": {"Japanese", "日本語", "ja", "jpn", "jpn", "jpn", "jpn"},
	"jv": {"Javanese", "Basa Jawa", "jv", "jav", "jav", "jav", "jav"},
	"kl": {"Kalaallisut", "Kalaallisut", "kl", "kal", "kal", "kal", "kal"},
	"kn": {"Kannada", "ಕನ್ನಡ", "kn", "kan", "kan", "kan", "kan"},
	"kr": {"Kanuri", "Kanuri", "kr", "kau", "kau", "kau", "kau"},
	"ks": {"Kashmiri", "كشميري", "ks", "kas", "kas", "kas", "kas"},
	"kk": {"Kazakh", "Қазақша", "kk", "kaz", "kaz", "kaz", "kaz"},
	"km": {"Khmer", "ភាសាខ្មែរ", "km", "khm", "khm", "khm", "khm"},
	"ki": {"Kikuyu", "Gĩkũyũ", "ki", "kik", "kik", "kik", "kik"},
	"rw": {"Kinyarwanda", "Kinyarwanda", "rw", "kin", "kin", "kin", "kin"},
	"ky": {"Kyrgyz", "Кыргызча", "ky", "kir", "kir", "kir", "kir"},
	"kv": {"Komi", "Коми", "kv", "kom", "kom", "kom", "kom"},
	"kg": {"Kongo", "Kongo", "kg", "kon", "kon", "kon", "kon"},
	"ko": {"Korean", "한국어", "ko", "kor", "kor", "kor", "kor"},
	"ku": {"Kurdish", "Kurdî", "ku", "kur", "kur", "kur", "kur"},
	"kj": {"Kwanyama", "Kuanyama", "kj", "kua", "kua", "kua", "kua"},
	"la": {"Latin", "Latina", "la", "lat", "lat", "lat", "lat"},
	"lb": {"Luxembourgish", "Lëtzebuergesch", "lb", "ltz", "ltz", "ltz", "ltz"},
	"lg": {"Ganda", "Luganda", "lg", "lug", "lug", "lug", "lug"},
	"li": {"Limburgish", "Limburgs", "li", "lim", "lim", "lim", "lim"},
	"ln": {"Lingala", "Lingála", "ln", "lin", "lin", "lin", "lin"},
	"lo": {"Lao", "ພາສາລາວ", "lo", "lao", "lao", "lao", "lao"},
	"lt": {"Lithuanian", "Lietuvių", "lt", "lit", "lit", "lit", "lit"},
	"lu": {"Luba-Katanga", "Tshiluba", "lu", "lub", "lub", "lub", "lub"},
	"lv": {"Latvian", "Latviešu", "lv", "lav", "lav", "lav", "lav"},
	"gv": {"Manx", "Gaelg", "gv", "glv", "glv", "glv", "glv"},
	"mk": {"Macedonian", "Македонски", "mk", "mkd", "mkd", "mac", "mkd"},
	"mg": {"Malagasy", "Malagasy", "mg", "mlg", "mlg", "mlg", "mlg"},
	"ms": {"Malay", "Bahasa Melayu", "ms", "msa", "msa", "may", "msa"},
	"ml": {"Malayalam", "മലയാളം", "ml", "mal", "mal", "mal", "mal"},
	"mt": {"Maltese", "Malti", "mt", "mlt", "mlt", "mlt", "mlt"},
	"mi": {"Māori", "Māori", "mi", "mri", "mri", "mao", "mri"},
	"mr": {"Marathi", "मराठी", "mr", "mar", "mar", "mar", "mar"},
	"mh": {"Marshallese", "Kajin M̧ajeļ", "mh", "mah", "mah", "mah", "mah"},
	"mn": {"Mongolian", "Монгол", "mn", "mon", "mon", "mon", "mon"},
	"na": {"Nauru", "Dorerin Naoero", "na", "nau", "nau", "nau", "nau"},
	"nv": {"Navajo", "Diné Bizaad", "nv", "nav", "nav", "nav", "nav"},
	"nd": {"Northern Ndebele", "isiNdebele", "nd", "nde", "nde", "nde", "nde"},
	"ne": {"Nepali", "नेपाली", "ne", "nep", "nep", "nep", "nep"},
	"ng": {"Ndonga", "Owambo", "ng", "ndo", "ndo", "ndo", "ndo"},
	"nb": {"Norwegian Bokmål", "Norsk (Bokmål)", "nb", "nob", "nob", "nob", "nob"},
	"nn": {"Norwegian Nynorsk", "Norsk (Nynorsk)", "nn", "nno", "nno", "nno", "nno"},
	"no": {"Norwegian", "Norsk", "no", "nor", "nor", "nor", "nor"},
	"ii": {"Nuosu", "ꆈꌠ꒿ Nuosuhxop", "ii", "iii", "iii", "iii", "iii"},
	"nr": {"Southern Ndebele", "isiNdebele", "nr", "nbl", "nbl", "nbl", "nbl"},
	"oc": {"Occitan", "Occitan", "oc", "oci", "oci", "oci", "oci"},
	"oj": {"Ojibwe", "ᐊᓂᔑᓈᐯᒧᐎᓐ", "oj", "oji", "oji", "oji", "oji"},
	"cu": {"Old Church Slavonic", "Словѣ́ньскъ", "cu", "chu", "chu", "chu", "chu"},
	"om": {"Oromo", "Afaan Oromoo", "om", "orm", "orm", "orm", "orm"},
	"or": {"Oriya", "ଓଡି଼ଆ", "or", "ori", "ori", "ori", "ori"},
	"os": {"Ossetian", "Ирон æвзаг", "os", "oss", "oss", "oss", "oss"},
	"pa": {"Panjabi", "ਪੰਜਾਬੀ", "pa", "pan", "pan", "pan", "pan"},
	"pi": {"Pāli", "पाऴि", "pi", "pli", "pli", "pli", "pli"},
	"fa": {"Persian", "فارسی", "fa", "fas", "fas", "per", "fas"},
	"pl": {"Polish", "Polski", "pl", "pol", "pol", "pol", "pol"},
	"ps": {"Pashto", "پښتو", "ps", "pus", "pus", "pus", "pus"},
	"pt": {"Portuguese", "Português", "pt", "por", "por", "por", "por"},
	"qu": {"Quechua", "Runa Simi", "qu", "que", "que", "que", "que"},
	"rm": {"Romansh", "Rumantsch", "rm", "roh", "roh", "roh", "roh"},
	"rn": {"Kirundi", "Kirundi", "rn", "run", "run", "run", "run"},
	"ro": {"Romanian", "Română", "ro", "ron", "ron", "rum", "ron"},
	"ru": {"Russian", "Русский", "ru", "rus", "rus", "rus", "rus"},
	"sa": {"Sanskrit", "संस्कृतम्", "sa", "san", "san", "san", "san"},
	"sc": {"Sardinian", "Sardu", "sc", "srd", "srd", "srd", "srd"},
	"sd": {"Sindhi", "سنڌي‎", "sd", "snd", "snd", "snd", "snd"},
	"se": {"Northern Sami", "Sámegiella", "se", "sme", "sme", "sme", "sme"},
	"sm": {"Samoan", "Gagana Sāmoa", "sm", "smo", "smo", "smo", "smo"},
	"sg": {"Sango", "Sängö", "sg", "sag", "sag", "sag", "sag"},
	"sr": {"Serbian", "Српски", "sr", "srp", "srp", "srp", "srp"},
	"gd": {"Gaelic", "Gàidhlig", "gd", "gla", "gla", "gla", "gla"},
	"sn": {"Shona", "ChiShona", "sn", "sna", "sna", "sna", "sna"},
	"si": {"Sinhala", "සිංහල", "si", "sin", "sin", "sin", "sin"},
	"sk": {"Slovak", "Slovenčina", "sk", "slk", "slk", "slo", "slk"},
	"sl": {"Slovene", "Slovenščina", "sl", "slv", "slv", "slv", "slv"},
	"so": {"Somali", "Soomaaliga", "so", "som", "som", "som", "som"},
	"st": {"Southern Sotho", "Sesotho", "st", "sot", "sot", "sot", "sot"},
	"es": {"Spanish", "Español", "es", "spa", "spa", "spa", "spa"},
	"su": {"Sundanese", "Basa Sunda", "su", "sun", "sun", "sun", "sun"},
	"sw": {"Swahili", "Kiswahili", "sw", "swa", "swa", "swa", "swa"},
	"ss": {"Swati", "SiSwati", "ss", "ssw", "ssw", "ssw", "ssw"},
	"sv": {"Swedish", "Svenska", "sv", "swe", "swe", "swe", "swe"},
	"ta": {"Tamil", "தமிழ்", "ta", "tam", "tam", "tam", "tam"},
	"te": {"Telugu", "తెలుగు", "te", "tel", "tel", "tel", "tel"},
	"tg": {"Tajik", "Тоҷикӣ", "tg", "tgk", "tgk", "tgk", "tgk"},
	"th": {"Thai", "ภาษาไทย", "th", "tha", "tha", "tha", "tha"},
	"ti": {"Tigrinya", "ትግርኛ", "ti", "tir", "tir", "tir", "tir"},
	"bo": {"Tibetan Standard", "བོད་ཡིག", "bo", "bod", "bod", "tib", "bod"},
	"tk": {"Turkmen", "Türkmençe", "tk", "tuk", "tuk", "tuk", "tuk"},
	"tl": {"Tagalog", "Tagalog", "tl", "tgl", "tgl", "tgl", "tgl"},
	"tn": {"Tswana", "Setswana", "tn", "tsn", "tsn", "tsn", "tsn"},
	"to": {"Tonga", "faka Tonga", "to", "ton", "ton", "ton", "ton"},
	"tr": {"Turkish", "Türkçe", "tr", "tur", "tur", "tur", "tur"},
	"ts": {"Tsonga", "Xitsonga", "ts", "tso", "tso", "tso", "tso"},
	"tt": {"Tatar", "Татарча", "tt", "tat", "tat", "tat", "tat"},
	"tw": {"Twi", "Twi", "tw", "twi", "twi", "twi", "twi"},
	"ty": {"Tahitian", "Reo Mā’ohi", "ty", "tah", "tah", "tah", "tah"},
	"ug": {"Uyghur", "ئۇيغۇرچه", "ug", "uig", "uig", "uig", "uig"},
	"uk": {"Ukrainian", "Українська", "uk", "ukr", "ukr", "ukr", "ukr"},
	"ur": {"Urdu", "اردو", "ur", "urd", "urd", "urd", "urd"},
	"uz": {"Uzbek", "O‘zbek", "uz", "uzb", "uzb", "uzb", "uzb"},
	"ve": {"Venda", "Tshivenḓa", "ve", "ven", "ven", "ven", "ven"},
	"vi": {"Vietnamese", "Tiếng Việt", "vi", "vie", "vie", "vie", "vie"},
	"vo": {"Volapük", "Volapük", "vo", "vol", "vol", "vol", "vol"},
	"wa": {"Walloon", "Walon", "wa", "wln", "wln", "wln", "wln"},
	"cy": {"Welsh", "Cymraeg", "cy", "cym", "cym", "wel", "cym"},
	"wo": {"Wolof", "Wolof", "wo", "wol", "wol", "wol", "wol"},
	"fy": {"Western Frisian", "Frysk", "fy", "fry", "fry", "fry", "fry"},
	"xh": {"Xhosa", "isiXhosa", "xh", "xho", "xho", "xho", "xho"},
	"yi": {"Yiddish", "ייִדיש", "yi", "yid", "yid", "yid", "yid"},
	"yo": {"Yoruba", "Yorùbá", "yo", "yor", "yor", "yor", "yor"},
	"za": {"Zhuang", "Cuengh", "za", "zha", "zha", "zha", "zha"},
	"zu": {"Zulu", "isiZulu", "zu", "zul", "zul", "zul", "zul"},
}
