package main

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	customFont := "arial-unicode-ms"
	customFontFile := "docs/assets/fonts/arial-unicode-ms.ttf"

	repository := config.NewRepository().
		AddUTF8Font(customFont, fontstyle.Normal, customFontFile).
		AddUTF8Font(customFont, fontstyle.Italic, customFontFile).
		AddUTF8Font(customFont, fontstyle.Bold, customFontFile).
		AddUTF8Font(customFont, fontstyle.BoldItalic, customFontFile)

	builder, err := config.NewBuilder().
		TryLoadRepository(repository)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).
		Build()

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	header, contents := getLanguageSample()

	m.AddRow(8,
		text.NewCol(4, header[0], props.Text{Style: fontstyle.Bold, Family: fontfamily.Arial, Align: align.Center}),
		text.NewCol(8, header[1], props.Text{Style: fontstyle.Bold, Family: fontfamily.Arial, Align: align.Center}),
	)

	grey := props.Color{200, 200, 200}
	for i, content := range contents {
		r := m.AddRow(5,
			text.NewCol(4, content[0], props.Text{Align: align.Center}),
			text.NewCol(8, content[1], props.Text{Align: align.Center}),
		)

		if i%2 == 0 {
			r.WithStyle(&props.Cell{
				BackgroundColor: &grey,
			})
		}
	}

	longText := "聲音之道㣲矣天地有自然之聲人聲有自然之節古之聖人得其節之自然者而為之依永和聲至於八音諧而神人和胥是道也文字之作無不講求音韻顧南北異" +
		"其風土古今殊其轉變喉舌唇齒清濁輕重之分辨在毫釐動多訛舛樊然淆混不可究極自西域梵僧定字母為三十六分五音以總天下之聲而翻切之學興儒者若司馬光鄭樵" +
		"皆宗之其法有音和類隔互用借聲類例不一後人苦其委曲繁重難以驟曉往往以類隔互用之切改從音和而終莫能得其原也我聖祖仁皇帝"

	m.AddRows(text.NewRow(10, "long text without spaces", props.Text{
		Top:    5,
		Style:  fontstyle.Bold,
		Family: fontfamily.Arial,
	}))

	m.AddRow(80,
		text.NewCol(4, longText, props.Text{Align: align.Center, BreakLineStrategy: breakline.DashStrategy}),
		text.NewCol(4, longText, props.Text{Align: align.Left, BreakLineStrategy: breakline.DashStrategy}),
		text.NewCol(4, longText, props.Text{Align: align.Right, BreakLineStrategy: breakline.DashStrategy}),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/customfontv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/customfontv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getLanguageSample() ([]string, [][]string) {
	header := []string{"Language", "Phrase: Talk is cheap. Show me the code."}

	contents := [][]string{}
	contents = append(contents, []string{"Africâner", "Praat is goedkoop. Wys my die kode."})
	contents = append(contents, []string{"Albanês", "Biseda është e lirë. Më trego kodin."})
	contents = append(contents, []string{"Alemão", "Reden ist billig. Zeig mir den Code."})
	contents = append(contents, []string{"Amárico", "ወሬ ርካሽ ነው ፡፡ ኮዱን አሳዩኝ ፡፡"})
	contents = append(contents, []string{"Árabe", "كلام رخيص. أرني الكود."})
	contents = append(contents, []string{"Armênio", "Խոսակցությունն էժան է: Showույց տվեք ինձ ծածկագիրը:"})
	contents = append(contents, []string{"Azerbaijano", "Danışıq ucuzdur. Kodu göstərin."})
	contents = append(contents, []string{"Basco", "Eztabaida merkea da. Erakutsi kodea."})
	contents = append(contents, []string{"Bengali", "টক সস্তা। আমাকে কোডটি দেখান"})
	contents = append(contents, []string{"Bielorusso", "Размовы танныя. Пакажыце мне код."})
	contents = append(contents, []string{"Birmanês", "ဟောပြောချက်ကစျေးပေါတယ် ကုဒ်ကိုပြပါ။"})
	contents = append(contents, []string{"Bósnio", "Govor je jeftin. Pokaži mi šifru."})
	contents = append(contents, []string{"Búlgaro", "Разговорите са евтини. Покажи ми кода."})
	contents = append(contents, []string{"Canarim", "ಮಾತುಕತೆ ಅಗ್ಗವಾಗಿದೆ. ನನಗೆ ಕೋಡ್ ತೋರಿಸಿ."})
	contents = append(contents, []string{"Catalão", "Parlar és barat. Mostra’m el codi."})
	contents = append(contents, []string{"Cazeque", "Сөйлесу арзан. Маған кодты көрсетіңіз."})
	contents = append(contents, []string{"Cebuano", "Barato ra ang sulti. Ipakita kanako ang code."})
	contents = append(contents, []string{"Chinês Simplificado", "谈话很便宜。给我看代码。"})
	contents = append(contents, []string{"Chinês Tradicional", "談話很便宜。給我看代碼。"})
	contents = append(contents, []string{"Cingalês", "කතාව ලාභයි. කේතය මට පෙන්වන්න."})
	contents = append(contents, []string{"Coreano", "토크는 싸다. 코드를 보여주세요."})
	contents = append(contents, []string{"Corso", "Parlà hè bonu. Mostrami u codice."})
	contents = append(contents, []string{"Croata", "Razgovor je jeftin. Pokaži mi šifru."})
	contents = append(contents, []string{"Curdo", "Axaftin erzan e. Kodê nîşanî min bidin."})
	contents = append(contents, []string{"Dinamarquês", "Tal er billig. Vis mig koden."})
	contents = append(contents, []string{"Eslovaco", "Hovor je lacný. Ukáž mi kód."})
	contents = append(contents, []string{"Esloveno", "Pogovor je poceni. Pokaži mi kodo."})
	contents = append(contents, []string{"Espanhol", "Hablar es barato. Enséñame el código."})
	contents = append(contents, []string{"Esperanto", "Babilado estas malmultekosta. Montru al mi la kodon."})
	contents = append(contents, []string{"Estoniano", "Rääkimine on odav. Näita mulle koodi."})
	contents = append(contents, []string{"Filipino", "Mura ang usapan. Ipakita sa akin ang code."})
	contents = append(contents, []string{"Finlandês", "Puhe on halpaa. Näytä koodi."})
	contents = append(contents, []string{"Francês", "Parler n'est pas cher. Montre-moi le code."})
	contents = append(contents, []string{"Frísio Ocidental", "Prate is goedkeap. Lit my de koade sjen."})
	contents = append(contents, []string{"Gaélico Escocês", "Tha còmhradh saor. Seall dhomh an còd."})
	contents = append(contents, []string{"Galego", "Falar é barato. Móstrame o código."})
	contents = append(contents, []string{"Galês", "Mae siarad yn rhad. Dangoswch y cod i mi."})
	contents = append(contents, []string{"Georgiano", "აუბარი იაფია. მაჩვენე კოდი."})
	contents = append(contents, []string{"Grego", "Η συζήτηση είναι φθηνή. Δείξε μου τον κωδικό."})
	contents = append(contents, []string{"Guzerate", "વાતો કરવી સસ્તી છે. મને કોડ બતાવો."})
	contents = append(contents, []string{"Haitiano", "Pale bon mache. Montre m kòd la."})
	contents = append(contents, []string{"Hauçá", "Magana tana da arha. Nuna min lambar."})
	contents = append(contents, []string{"Havaiano", "Kūʻai ke kamaʻilio. E hōʻike mai iaʻu i ke pāʻālua."})
	contents = append(contents, []string{"Hebraico", "הדיבורים זולים. הראה לי את הקוד."})
	contents = append(contents, []string{"Híndi", "बोलना आसान है। मुझे कोड दिखाओ।"})
	contents = append(contents, []string{"Hmong", "Kev hais lus yog pheej yig. Qhia kuv cov code."})
	contents = append(contents, []string{"Holandês", "Praten is goedkoop. Laat me de code zien."})
	contents = append(contents, []string{"Húngaro", "Beszélni olcsó. Mutasd meg a kódot."})
	contents = append(contents, []string{"Igbo", "Okwu dị ọnụ ala. Gosi m koodu."})
	contents = append(contents, []string{"Lídiche", "רעדן איז ביליק. ווייַזן מיר דעם קאָד."})
	contents = append(contents, []string{"Indonésio", "Berbicara itu murah. Tunjukkan kodenya."})
	contents = append(contents, []string{"Inglês", "Talk is cheap. Show me the code."})
	contents = append(contents, []string{"Iorubá", "Ọrọ jẹ olowo poku. Fi koodu naa han mi."})
	contents = append(contents, []string{"Irlandês", "Tá caint saor. Taispeáin dom an cód."})
	contents = append(contents, []string{"Islandês", "Tal er ódýrt. Sýndu mér kóðann."})
	contents = append(contents, []string{"Italiano", "Parlare è economico. Mostrami il codice."})
	contents = append(contents, []string{"Japonês", "口で言うだけなら簡単です。コードを見せてください。"})
	contents = append(contents, []string{"Javanês", "Omongan iku murah. Tampilake kode kasebut."})
	contents = append(contents, []string{"Khmer", "ការនិយាយគឺថោក។ បង្ហាញលេខកូដមកខ្ញុំ"})
	contents = append(contents, []string{"Laosiano", "ການສົນທະນາແມ່ນລາຄາຖືກ. ສະແດງລະຫັດໃຫ້ຂ້ອຍ."})
	contents = append(contents, []string{"Latim", "Disputatio vilis est. Ostende mihi codice."})
	contents = append(contents, []string{"Letão", "Saruna ir lēta. Parādiet man kodu."})
	contents = append(contents, []string{"Lituano", "Kalbėti pigu. Parodyk man kodą."})
	contents = append(contents, []string{"Luxemburguês", "Schwätzen ass bëlleg. Weist mir de Code."})
	contents = append(contents, []string{"Macedônio", "Зборувањето е ефтино. Покажи ми го кодот."})
	contents = append(contents, []string{"Malaiala", "സംസാരം വിലകുറഞ്ഞതാണ്. എനിക്ക് കോഡ് കാണിക്കുക."})
	contents = append(contents, []string{"Malaio", "Perbincangan murah. Tunjukkan kod saya."})
	contents = append(contents, []string{"Malgaxe", "Mora ny resaka. Asehoy ahy ny kaody."})
	contents = append(contents, []string{"Maltês", "It-taħdita hija rħisa. Urini l-kodiċi."})
	contents = append(contents, []string{"Maori", "He iti te korero. Whakaatuhia mai te tohu."})
	contents = append(contents, []string{"Marati", "चर्चा स्वस्त आहे. मला कोड दाखवा."})
	contents = append(contents, []string{"Mongol", "Яриа хямд. Надад кодоо харуул."})
	contents = append(contents, []string{"Nepalês", "कुरा सस्तो छ। मलाई कोड देखाउनुहोस्।"})
	contents = append(contents, []string{"Nianja", "Kulankhula ndikotsika mtengo. Ndiwonetseni nambala"})
	contents = append(contents, []string{"Norueguês", "Snakk er billig. Vis meg koden."})
	contents = append(contents, []string{"Oriá", "କଥାବାର୍ତ୍ତା ଶସ୍ତା ଅଟେ | ମୋତେ କୋଡ୍ ଦେଖାନ୍ତୁ |"})
	contents = append(contents, []string{"Panjabi", "ਗੱਲ ਸਸਤਾ ਹੈ. ਮੈਨੂੰ ਕੋਡ ਦਿਖਾਓ."})
	contents = append(contents, []string{"Pashto", "خبرې ارزانه دي. ما ته کوډ وښایاست"})
	contents = append(contents, []string{"Persa", "بحث ارزان است. کد را به من نشان دهید"})
	contents = append(contents, []string{"Polonês", "Rozmowa jest tania. Pokaż mi kod."})
	contents = append(contents, []string{"Português", "Falar é fácil. Mostre-me o código."})
	contents = append(contents, []string{"Quiniaruanda", "Ibiganiro birahendutse. Nyereka kode."})
	contents = append(contents, []string{"Quirguiz", "Сүйлөшүү арзан. Мага кодду көрсөтүңүз."})
	contents = append(contents, []string{"Romeno", "Vorbirea este ieftină. Arată-mi codul."})
	contents = append(contents, []string{"Russo", "Обсуждение дешево. Покажи мне код."})
	contents = append(contents, []string{"Samoano", "E taugofie talanoaga. Faʻaali mai le code."})
	contents = append(contents, []string{"Sérvio", "Причање је јефтино. Покажи ми шифру."})
	contents = append(contents, []string{"Sindi", "ڳالهه سستا آهي. مونکي ڪوڊ ڏيکاريو."})
	contents = append(contents, []string{"Somali", "Hadalku waa jaban yahay. I tus lambarka."})
	contents = append(contents, []string{"Soto do Sul", "Puo e theko e tlase. Mpontshe khoutu."})
	contents = append(contents, []string{"Suaíli", "Mazungumzo ni ya bei rahisi. Nionyeshe nambari."})
	contents = append(contents, []string{"Sueco", "Prat är billigt. Visa mig koden."})
	contents = append(contents, []string{"Sundanês", "Omongan mirah. Tunjukkeun kode na."})
	contents = append(contents, []string{"Tadjique", "Сӯҳбат арзон аст. Рамзро ба ман нишон диҳед."})
	contents = append(contents, []string{"Tailandês", "พูดคุยราคาถูก แสดงรหัส"})
	contents = append(contents, []string{"Tâmil", "பேச்சு மலிவானது. குறியீட்டை எனக்குக் காட்டு."})
	contents = append(contents, []string{"Tártaro", "Сөйләшү арзан. Миңа код күрсәтегез."})
	contents = append(contents, []string{"Tcheco", "Mluvení je levné. Ukaž mi kód."})
	contents = append(contents, []string{"Télugo", "చర్చ చౌకగా ఉంటుంది. నాకు కోడ్ చూపించు."})
	contents = append(contents, []string{"Turco", "Konuşma ucuz. Bana kodu göster."})
	contents = append(contents, []string{"Turcomeno", "Gepleşik arzan. Kody görkez"})
	contents = append(contents, []string{"Ucraniano", "Розмова дешева. Покажи мені код."})
	contents = append(contents, []string{"Uigur", "پاراڭ ئەرزان. ماڭا كودنى كۆرسەت."})
	contents = append(contents, []string{"Urdu", "بات گھٹیا ہے. مجھے کوڈ دکھائیں۔"})
	contents = append(contents, []string{"Uzbeque", "Gapirish arzon. Menga kodni ko'rsating."})
	contents = append(contents, []string{"Vietnamita", "Nói chuyện là rẻ. Cho tôi xem mã."})
	contents = append(contents, []string{"Xhosa", "Ukuthetha akubizi. Ndibonise ikhowudi."})
	contents = append(contents, []string{"Xona", "Kutaura kwakachipa. Ndiratidze kodhi."})
	contents = append(contents, []string{"Zulu", "Ukukhuluma kushibhile. Ngikhombise ikhodi."})

	return header, contents
}
