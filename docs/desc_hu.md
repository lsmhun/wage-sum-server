# Kitekintés a Go programozási nyelvre

## Bevezető

Az év során alkalmam nyílt részt venni egy képzésen, ami bevezetőt
nyújtott a [Go programozási nyelv](https://en.wikipedia.org/wiki/Go_(programming_language))
alapjaiba. Az ilyen kurzusok ugyan hasznosak, 
de ha csak ismerkedni szeretnénk, akkor a legjobb ötlet a hivatalos
[Tour of Go](https://go.dev/tour/welcome/1) online anyagon végigmenni.

A következő szint, hogy ha valamit össze is akarunk rakni az új eszközünkkel,
akkor neki kell látni valaminek, amiben ki lehet próbálni a _helloworld_ 
példáknál nehezebbet és egyben a toolingot is.

Erre a célra egy bértömeg számoló alkalmazást választottam.
Arra képes, hogy a dolgozói hierarchiában megmondja, hogy
adott menedzser alatti dolgozóknak mennyi az összes bérköltsége.

![Wagesum application](wagesum-app01.png)

Amint látható, egy egyszerű CRUD alkalmazás, ami REST interface keresztül szolgálja
ki a kéréseket, adatbázisban tárol adatokat és goroutinokkal segít javítani
a sok lekérésből adódó terhelést.

Szubjektív szempontokat fogalmazok meg, és a változtatás joga is megillet. :)
Állítólag a nyelv ötlete is akkor fogant meg az alkotók fejében, amikor 
a Google-nál éppen kávéztak egy hosszabb C++ buildre várva.


## Kezdeti élmények, vélemény

Nekem olyan élmény, mintha a C nyelvet Python tulajdonságokkal, fél szemmel
folyamatosan a Java-ra és Spring-re tekintve készítették volna, figyelembe
véve és lehetőség szerint kerülve a C++ problémáit. A világon 
nagyon sok C és C++ fejlesztő van, akiknek a nyelv kényelmes
lehetőséget kínál a microservice világba való átlépésre.

Vannak pointerek és van _nil_ is, viszont nincs pointer aritmetika és nincs
operator overloading sem. Nem hullatok könnyeket értük. Van Garbage collection, 
de alapból belefordul a binárisba, nem külső JVM mechanizmusok vezérlik.

Statikus linkelést használ. Ennek eredményeként képes sok mindent egyszerűsíteni,
illetve ennek is köszönhető a rendkívül gyors fordítás, valamint az elérhető
egy (!) futtatható állomány, ami mindent tartalmaz. Van ennek olyan
hatása is, hogy a függőségek kezelésénél képes kiszűrni a nem használt 
library-kat, egy sima _go mod tidy_ parancs helyrerázza a felduzzadt go.mod leírót.

A kód egységes kinézetéért beépített formázó és elemző áll rendelkezésre.
Azaz könnyen és fájdalommentesen lehet egységes kód képet elérni mindenkinél.
A hibakezelés kicsit furcsa a modern nyelvek kivétel kezeléséhez képest,
de meg lehet szokni. 

A Go nyelv nem objektum orientált nyelv, de _struct_ segítségvel
elérhető hasonló működés, leginkább a Pythonra hasonlít, ahogy 
_func (p MyStruct) f(p param) (RetVal, error){...}_ ránézésre is mutatja a rokonságot.

A funkcionális paradigmát viszont messzemenően támogatja. Nagyon könnyű
és kézenfekvő a kód közben is function-t definiálni. A wagesum alkalmazásban
a [emp_sal_service_test](../internal/pkg/emp_sal_service/emp_sal_service_test.go)
él is vele, de egyáltalán nem használja ki a teljes potenciálját.

Mindeközben a C őst nem is akarja letagadni, az interoperábilitás 
biztosított. [Ebben](https://programmer.ink/think/interoperability-between-go-and-c-language.html)
a cikkben élő példát mutatnak be. Van azért hátulütője is, az SQLite 
támogatásnál például rögtön _gcc_ után vágyakozott, így végül ki is vettem,
pedig jó lenne egy pehelysúlyú, in-memory adatbázis a teszteknél.

A [Convention over Configuration](https://en.wikipedia.org/wiki/Convention_over_configuration)
elvet talán kicsit túltolták. Az még tök jó,
hogy a _private_ és _public_ kiírása helyett a láthatóságot a kis és nagybetűs
kezdés jelzi. De amikor comment-be lehet írni test assertet, akkor azért 
felszalad a szemöldök. Oda azért nem kéne üzleti logikát írni.


## Openapi és kód generálás
A kedvenc openapi generálók mind ismerik a Go nyelvet. Így most is elég volt
egy tisztességes API leírást megalkotni.

```shell
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api/wagesum-openapi.yaml  -g go-server  -o /local/internal/pkg --additional-properties=noservice,enumClassPrefix=true,featureCORS=false,onlyInterfaces,outputAsLibrary=true,sourceFolder=openapi
```

Így teljesíthetjük az _API first class citizen_ elképzelést, és még kódot se kell
írni. Viszont a generátor nem a legfrissebb, ötleteket azért így is ad, de a linter
persze már tiltakozik. 

A wagesum alkalmazásban az _internal/pkg/openapi_ alatt találhatóak a generált 
források, de a _service_-eket átmozgattam másik package-be. Ez azért volt hasznos,
hogy elkerülhessem a körkörös hivatkozást.

## Eszközök

Egy programozási nyelvnél manapság kevésbé számítanak a tulajdonságai,
ha nincs hozzá megfelelő eszközkészlet. A Go itt nagyon szépen teljesít,
alapértelmezetten is rengeteg kényelmes eszközt kínál.

A legtöbb eszköz nyílt forráskódú, barátságos licenszeléssel. Ha Java 
világból érkezünk, akkor a Hibernate helyett ott a [gorm](https://gorm.io/index.html).
REST hívásokra van [gorilla mux](https://github.com/gorilla/mux), 
létezik az [echo](https://echo.labstack.com/), de alapszinten meg is lehet
írni a beépített eszközökkel. 

Project felépítésének kialakítására [ezt](https://github.com/golang-standards/project-layout) 
a layout formát találtam, de sokan használják a _/src_ könyvtárat a forráskódoknak.
Mint látható, próbáltam tartani a szabványt.

A [Viper](https://github.com/spf13/viper)t a túl sok függősége miatt 
dobtam, mert nem volt szükségem a lehetőségeinek többségére. 

Létezik egy rakat féle tesztelő framework, de nem annyira triviális 
velük dolgozni, mint mondjuk a JUnit/Mockito párossal. De a tesztelés
már egyből támogatott és azonnal rendelkezésre áll, bármilyen
függőség nélkül. Tehát fapadosan azonnal lehet akar TDD-vel indulni,
még akkor is, ha később mást is hozzá szeretnénk adni.

## Go rutinok és alkalmazásuk
A Go rutinok az egyik legelegánsabb és legérdekesebb megoldása a nyelvnek.
Ezek lényegében pehelysúlyú szálak, amik az operációs rendszer szálaitól 
függetlenül üzemelnek. 

No persze, nem az első, lehet itt Erlang-ot is emlegetni, de a Scala-nak
a [Zio framework](https://zio.dev/) szintén kínálja ezt, csak ők _fiber_ néven
nevezik. A Java-ba a [Project Loom](https://openjdk.org/projects/loom/) 
fogja elhozni nemsokára. 

A wagesum alkalmazásban a lényeges rész az 
[emp_sal_service](../internal/pkg/emp_sal_service/emp_sal_service.go) 
goroutinjai teszik ki. Az ötlet annyi csak, hogy a hosszan futó, rekurzív
lekérdezéseket párhuzamosítsuk, és _channel_-eken keresztül dolgozzuk fel.


## Összefoglaló

A Go nyelv egy jól összerakott, jól támogatott programozási nyelv,
kitűnő eszközkészlettel. Alacsony memória és processzor használat,
hatékony megoldások jellemzőek rá. Számos területen viszont évtizedes
tapasztalatokat hagy figyelmen kívül. 

A Java világából érkezve nem érzem azt a késztetést, hogy 
azonnal áttérjek rá. Ott is vannak előremutató kezdeményezések, 
Scala, Kotlin és mások. Nyilván sokkal gyorsabban produktív tudok lenni. 

Azonban megvan a helye a nyelvnek a palettán, a Kubernetes világában
széles körben használják, így sikerre van ítélve. Mindenképpen megéri
kipróbálni és kicsit beleszokni.

## Linkek
* https://towardsdev.com/golang-productivity-hacks-part-3-auto-generating-test-4c8055dc7946
* https://eli.thegreenplace.net/2021/rest-servers-in-go-part-4-using-openapi-and-swagger/
* https://stackoverflow.com/questions/7106012/download-a-single-folder-or-directory-from-a-github-repo
* https://ribice.medium.com/serve-swaggerui-within-your-golang-application-5486748a5ed4
* https://github.com/GoogleCloudPlatform/golang-samples
* https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
* https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266
* https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/ 
* https://swagger.io/docs/specification/authentication/oauth2/
