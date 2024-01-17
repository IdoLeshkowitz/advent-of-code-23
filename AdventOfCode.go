package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var input = `
eightqrssm9httwogqshfxninepnfrppfzhsc
one111jxlmc7tvklrmhdpsix
bptwone4sixzzppg
ninezfzseveneight5kjrjvtfjqt5nineone
58kk
5b32
1dtwo
six7two7sixtwo78
mvhsixpptztjh13sixthree2
six1bqqvrxndt
fourmk5grmqone944nbvtj
twofiveqxfivezpkvfvxt5eightjhnpl
fpfqp7three7
scmlf76ninegjzjkj97two
fivetkhfnnx22
sevenxvbcbsvxr7eighttwo
1hvhqqmrs1bgttshthg6
4bvnccbdh4onefztdrpq62vvbnvpxxvgrngnfjgfk
653spgrvd
sixctlhkjmmxh2fourfivenine37
229mjp3txmqsxxqdbnnnbrtrcctgzseven
jfourdbpcjc39bhglgnine
bvnltxdmsp7twoxzpdjdvkxeight4twothree
jlvcdrkhzh8seven3
418oneeight
53flcrlvqdeight84frmdcsixchcbc
114sixone1eight2
xrbtzbklqsl11
bhfhszrhzgrhsfd2threeseventwosevenoneseven
four9one
5p
twovhjpdxmcxshnhv5vs
qkkqeightcxcltnn7one9pmhlmvsxnine
4cbptmvp1
84xgm
bzsmqhkrdtdmhhjgrjsdfour1ninetwo61
onetwoeightgflhlgksevennine7two6
mbjhkhfour6
8cvqk6eightonethree1
qhbllbnlkr3rcsmjvztgd
18eight4
hhc6onegvkgkqs5mvsone
66bnfj
one99xvrhninefive
eight96nxcjjddmseightxvgsixfiverrzpvmgnl
rpgpczdsxpjgql39
855dnthhxld6eight
four29twosspz1
sixfiveqvrbcdr9fourlrkpkmxphlsbone
341
mhqjjg9six9nine
7pvjctsgvsix64
75twotwothreegcvssgbvhpzcnbgteight
7keight8eight
52threerhfmklssxcptmnlr4hqc4
xndfqvgxn3five
974lknineseven
rlnsix3
771m1
xvtjhq7six64threeeightgspmxgv
4sbqdxbmmzj6fiveone
onesvvch4rvhmvncnk
mkzsftp69six6
fiveptnn7
94lsgsjxrrghxxsr4
1grnvgpeightjthqmrfnszpfhfninefive9mbtf
qtsdfour2
5mfhmskpcvqbxjzzxt4lq3sevenkv
rbhjk1cdzjhtzkcbtvmfm
nine4eightpmrptkb
bfiveeight1lxzkzvbtkkgxxs38
mxcgbjqvhd1sevensevenrgp7two
fourxrdzzmjfmtr62one
sevenppqtlhvtwo7phlrbssxb
dhbnjmxg3bsgbhmlfiveseventwo
twoonemrbftgtzeightqjmjctmq55
ggk2lt586dfzqbjsvj1one
jtgpzjjtwo86seventwo
37nine4onebqvsnmvg277
hmpnzmqsfour6
xzfhgzllmcbc56vpbpbbjffmgr3jrc
5lclone
hjbvkdtmrgvpfive9sevenfive1nlzqlkfrg
htwonetgxvjdkrvjsfjjbfmcthseven29
six89bdlssd
3eight6one
vtbsix2twolzrhfr1
vhdcvtj5
167nsnmgxhtvn54fivedcbgrhm
4three1five82four
15eightonethreesixthree
dvhtsccljt51
pbnfrxblk3sevenxjcmcvhlgrghpbgdnpl8xsr3fiveoneightq
242three8
2fivehgqfxgl8kpdknxhmk5bmmsbz
fivejvjeight6fsqgtpvcb
1threelgkbhlhhlmrbvxqqgf
klmqfgfg1gnine253psn8
47eight
eight83mvdtsqppjhgjnsvngfive
7mcmzvsv6seven
56seven98three4three
pfnbthreegthreefknjm4five
2nqgrdcshfpjfpqdrvnq1twoccpmxpxvv
xtwo7threemxbtpsvjkgrfivethree2
9pkdfourfour1zjvczkhpbj
1three2
pkkbphkgqfivellrnvnkdxpql3
ppc62
one73ptfxsbbpqqgctdjhzjsjc91
nine7threefourvvk
six59542xcxqcbnrvzfbshcxxddz2
ftsfj2ninesix1hdjsrpkonelklfpltv
ninehmxgkqbmhvtlvdmdtvpeighttqfour1three
frbnineeight4168ksmjstpqvzhnn
rgnrntwohvqhgxxfkonefour4mfdr5ftgtjjv
3ninejbszdvdgznfourxpcxspqxnthlngkncvnineccq
threemjglxtp5cqmtwotwo2seventwonerkl
eightsevensevenlmbjzprggthree1eight
57nineninezdcf
cqoneighteightjnrfkplvninefivemck18mnhszhkv4
tbvdcsjsvmxtshv3fourseven4kmxvvfour9
bxcsix19six8dnqsbx
7five81ncchkdk
four4ck7rtjmjpccpeightone
fivetwo6nine1tdczktmfninelrbnnine
onetwo9twoeight5sevensix
cvvtbmninebneightsix1dnnfkgmnm
h7three3
tpnzsdm9sixtg5sixqvcqsq3
1sevenzmbcpgtfrjvq
r8757
fournineseven6fourfour
798dpbrkfourtwoxdrgqkrkmfeight7
threebvqqjcldjx4nine5
3hbl
twodndcfddkvfivermvkrfzsnqthree5
mhdcvsixmnqlvmvxmxfour3ct
25dpfsrbcllhtwothree2pthreezfhjx
dcfggnine1onetwoone
vtbmbpgffive2hdmzjzqqqc4one
zg11
txrknhvhbv183
dlvscqszz82nvtpb7tktvtgjbml
twofour5sevensmqfjrjcndmvcvqdfsrsix
446sixeight6rbrltdzf
132ncq1
4ninezdfzgvzf4four7qkzstcq
7rx5xntgxfpmvsevenzmzmbjlc3fivefive
threeseven7tshthree
gnbqhninecjnhlpcfivenine18
gjntwonenllmzgqsvq36lc45fourdrtzlctr
3tqgbfrk
zldl3zxpfbpveight
2ninetgppcvqrq
7one1sixeighttxcnhltwooneeight
bsm3hslqcr8xslndqnnvfpzvprlkt
khvptwo7kkbznndpqsevensevenvlr
sevenscneight8one4qnkc
2v
2threerjnineonev
68qhknonebtxvmqh
3btb
kcxbqzbjqt3twofourfxdlprsxkzlmflbveight
9bmdcninecjdv7
ceightwoninelkbbfxgsv9fb5n
ztwo2
7four9cpkclqxtrmpdgzxgtwo
8dghrmgprdjeight8knnb852one
2three36eightfournone2
6ztwofkzlhvjdrxtsmlbgczf
knine6ninebpmzjbkg9tttkhtgcklbfive
sevenv3
411one4483
88jvjggxqfour3zrbvbxjzmthree
fivelqcnqfvgp18t
fivejceightqlsdrmrnbzfbjskstzrllxrdlcxpjkvf1
832
sixfive77rhkjdhvbpdfjxpkmfdq66qqtfpfs
fv6svkbnsgtpznblnvkvk5
sixsevenfoursixeight486
gsbfrjpngshpmlxf2
gbmmmvdhffbbcq3
5fivehxx28
snjxttwo1zd5sixklfl
szsvdzsix3nine32nine
onervkhknmnsix3four363
1btphrrvxdeightonekdhv8
gkphmq73lfhflk66xpfive2
611four3gnjsdkvksjdxfiveeight
klbvb9zk4eightninetwo
mreight59sevengqbhnspvhqcj2six
hnvgkmljlpthsgjrzmsevensix6fcvtxddbnx5
ckfvkblhvv6gbsnlsevenktblt29three
eightfoursevendnsghkfsg8fivextnnine
5hrdqmfjq
twojcvkkjklzmfive2fourxqgdsdgzrjltwo7
onetwonvxnxnntxcthree3
threentwonine6
sprfmxlqvb9jnbzltskxppqzdscrvhpfvjjcqhxcf
sixhfour95
41fourqhpjbztknqnfpxvzc
two23sevenfgmmnszone113
one41seven57
oneninebqqklhk6gmdzddbhgzqcmxxfnmrvr
4mkpgv87eightckzjjfm
vhgdmvncxn88ldbt7464
gcczfprplf7
6eightninesxthreefive99
two69fivecjxjhgjmgvttl9
mztbzjmgxnpkfrqnrbgkgfourjjfjtone8eight
63mggldkcprlz
6sevenfourfive
1zvmrdnpzcsqqmxscn
five5tkgb8rrztmcfivebknjd
7pscpfive
eight9sixgfvhvlcnineeight
seven99fourtwojvxfrqmrdlthree
sppcgnjzeight93j
fivek7seventhree
twornbhtrlnznpctrxhqtngzdtnvfb2
sevenfivefour63five38nlrxfcjpm
qvtdcspc4zxmmhpklhcdlznqfg46lct5one
1bnndtnsfjdsevenfivetwo3k85
8five9six
3ltcc7trmmhbbbpjfive
cmbchhhczmqlp3
nctz78twodljcqvplcqg
nine4sevenpnbbztpvkbgztb
zseven9eight
fourktzscmnrvddnnzsqfnfctzpdvtwo9
nrtbgdjpm2ldcfdm5jjhx9
jq9two68kjttwo67
lnneightfourzqz6lgvxnthreeseven7
fiveglp85
ninefouronesix7k1three
3five5sevendhtmjhbh2
rbjhnmmgsvmtk8four
vtrvvjsixhnctwocvskgzt3two3
kxfive5threezgtd2
2fivemcnngtzxsgbxmvbl
eightlrlztkvhfivefour5
2twojcq7qrrbddmpsb
3twosixthreebqtoneone2
nine9six3vlmpqbgjjqdftldpq
tgrglqfxxc2onetwo76oneonex
one48one
fivesgsnrzsms6one
fiveseven1tzhxdknkseven4
3sprtonefivelxg47
twotwoseven3ninenqdvxgm
nskjpvvqeightnine93fivecngkjcd3
ltfxscllxk9pjznpnmhfmrzmqbq
4nineeight6four7
fjtpj763
four6six73htbstbbpztwofj
oneonetwo2five43three
five8bgcjnlzcgqpfkn
11mcn
9twovhkltdpkqzc65six
8fkprfkg9xfjxspqpshlfkqpnrt
4sblrf7745
lblxmbzfour6187tggqllj
dhdmlx71mbbxtvhszhsvcm
six1qbqglfsssx
164tsthbb2
8nineldrtltqfivebqnrthrm
sevenlsjbsklhxxfiveclbldxxtrdllxzthree6eight
18gnkrxfmnineone
six4threetwoeightgcn
vmknclnmnphb2czdbjmcone
eightq67
eighthnzslhbblr85eight3
3457kdzhnppqz5four8two
crcskvmhthree41
8pccnsbv8ln3
4eightonevd
pdz9threenine
3sg5gkrncz
6dxnslkl3xqlnm965twonexxn
9nlhmmkzsdbpdctd7ninec
62eightnine7nine3lrd
qbprlzczreight7threegqnrshrhc
3qlmr
three1twojgptkzgxmf
twosvdsfourone8
5sixfourzvjtkpk
six418
five523fivecbs
nineninesevenztfggvfkgkzfcm2
qnsix5dnv7three
three5twofour
hlmtnzsmlnjxdtwo6
9sixnglrctg
onecrgfq5hdldpc
nhzctlx94eight
19djvld
tzp8zzv8six1
1nsnine5sixhqxfk
33sixtwojrdvksrfsnltglggxdhbsxf
bbvtpxptrnvjqzl3rldseventhree6
2lcntfphb2lgpjbdeight
fivetwo1jjgkt9kltwo1nc
7xnmrscpfkthreejqbhlrtf4sixrbfrone
3threetqfkv1twofive
95ninejlftlxrs1nxxfsqz
hrsrszgrcl9seven8eightksdnhqsq7eight
5436ninefour
one5five
prcmxone8lhkblvr714three
92btgsllktgf7fivejhgsg2eight5
fourmm61nine558nine
bxjx2
one6onetwotjxthreefour8
threedkpnpfvgt1one3nine1
bxfour3two2sb4twondmfdpsz
11sevennrpxftwooneeightmx6
9mqxcrjxnp7hdjgqktxm
2rnjlg7mbxstzbdh
qmsixhckzone1
qzsnq6sixtwosixtwohhgbsrqgnine3
foureight7scksqtkmnfiveseven765
15oneckvshqd
41tzlxsfivebsckffdps
threefivefournine7
5fiveonefour8lhqmltwoeighttwo
four5six
mhnrspfourmflmvkc52cjkvxheightsevendtddjdcnb
vxrrlfnlqf1twoeightninesixonetwo
8xonetwo2
rveightwo79three
onebjlr9sixldqrbtwo4
f1lhrbsix
mbbkv7ffpk
8twofourmxqvkqfcjfoureightplgpmrtxm3
2rtrxjzqeighteightqtmsfnpdscpgqvxd7
nineqggljvzvxltwozsvsfournine9
ninexpmnnvqsfhnprqrqlcgfive9mtnflvttwoqlgphhb
11kbpmv1
htspdnh1xhbbh3lzcjjx1
five6fivefive5six
sxtbktj7
ncvkgvgbeightfour89ttbrjthree
xcgxzxbfnkxdqn73eight
5nine4fivesixtwotwo1
ffgzdfhn6
fourzvtfcczlxhnnx5three
threednfntx4eightwovql
sixtwoveightnine7twonineseven
seveneightcclmgknrgninemnjsrsqsevenfftmlqkch7
pvqnltjs2hghkrphnine3scngkjfcsn
62vdnbzrcrjsndqqr2
5cxhscqpgdzbrnnq1m9b
gnkclhmbjfourfivemmxpqx2qlxvsix
sevenxcp4fourdlqgpxcl
8eighttwoone
bjslbfrspcnffnine9rvnjjrvcsix2
six7tfive6hkllf56six
km3
sixtjt2threefour2vqqcxj84
9vflltwo4five
nineprprrcjt3eightmxfour
4twoninehvsbszqr
dtdmkcsd41eightfourxppqzkjb
xmdmghzdp9sevennine94
2bbjsdlxoneeight
ninefivetwonine8
68four
ninepdjpfmzxthree3dkxgttvncbr
52threesevenninefour6lfrlrsgzk
ghkczjt86bdk3czvhcone
cdjsd6jhnnnhzbzllqdjgpgnninevmcvbcxxltsix
1mxfsrninegfmgvnine14hbfnshgbprone
39four8
bxnnjqkninetwotwo58txgvrbxvq
k4bftq68seven4nineseven
seven97fivekxjnseven4fourfour
drgttpqpsevenvrkxdlmvtctsc72seven
eight298
2zsqmjskp
six5bgdkhzqt
fournlknxg35vqdqmnln6f
565rqtzdpqhlldxgnine7oneeight
3jfthree7vlkpfour1
fourgngnqtgd675vgrrjf
25nrfive
kmlbnrm5dtvqndldh
vgbzkpnltxrp5tpvb
5mbzzk3nine7cqkngz1mm
sevenfive952
4pxk8four
seven87fourlzlnineone
zxllsctgqmsevensix72
9ggkqvsrhftwobndlt
8sixsfzlfpztjtwofourqvnptkgllxcf
5dvsjvtm
29dsvjrl3pvgjqncbgcxc
9mhvrb8fbtppbhm2s
eightninel5
jltnzpcdr8one5szgf3nine88
8xzgs4chdkfour
three7vgnbtqvhthree8sixq
2onesix354jj58
175rpdmxfeightwos
gmngst7hpfvgmtfrqbb
9hd2lsxprdvtqxcv55
four13cjqkvgxvbseven8
1kqfrqsevenqgjttjrspd
7seveneight6zfmdbzdj2
9sixpfjbnthreedgbhblmr
7hdqqqkone6htzthree
hmnxstkbzlhqjpdn3three2
dpbjgmnn1
ninefivecznsbttpcvkthreeoneeight1zqtxb
6hlxdlpgrl7six
fourdhczrzd9mmmdthcngsmdqkq9fivetwofkb
7fbcrzcxjvxtqbfive168
sevenhlrkxgrggkqgd12
vzmpvhqvkpdfmthreetwo9
21cqxtjtwoljsixxbf3
bzzxkxtl3rhsrpnnzseven
threevhdqqvtwothreenine48frqsfhgsgptbbn
foureighteight3pgbcftjdbbsmcqjcrmt21
twovpfbsd42five
6foureightwofh
5185cbhgvkvpfzk1
1one3two5
seven3mdjks6kctnnchjgpnineeight
sixfivecjfkx5
sccbfqfive28bhconexmztpcftrbz
qfzhgl1nine5four59nbxhclpk
mjrvgdz3nine6gkvznv83
mfxqslvpzeight2
ztwoneeightknnjh4nine
onefivesevenfsmmhkbcplj6seven
2lfgzblvdflgxnsqfxtksbb
blmvzczjs61fourmtds
5two2
jqppjfxfour2
3five2bcrn653
86five6bseven
rzztlcbvteight135
1sjngcngjrsht3ninehninefive
81sixkrhvrhxqhn65
dkrgmnlcbjdjxblbfnp5
three85xrckdqtjqphjsjqflmt
5fmptvmz19fourmbzrttnxnk
dptwonefbqhrxtljddtkhh6four23
xfvgkfive69
cbtfrbpxphj2sevenmmzrpccnine
5twotwoseven
fivetwoeightonefive954
zxccrkvgdqtklbnhtdtdsghcseventwo6
three6vqmtchfdjxveightone
onenine39twodksvrdsxflthree
threesbpseven5zqtwobtmpgqjg1
326dgjmzzfivehkhsjrseven
dgthreeeightthreefbhbltbdjnineseven3klvx
4xkcqeighttwo2
eightonerhlnchp17
2jcmzbczstsrmbpzxpftkmznzckhv7three
7foursixbrcc6twosixgnf
4lqxfourtvxhqtlhlx7xcfxhmqzbone
2sevenz6sixtwo
nine1pv
95jzlkxn
onerzfnqhmtjpqff93lrctjgqhseveneight
5eightrdjnine3
threetwo6fourrcrq9dfmbsznshkfqmpvcb
five3hpjznhbtjonecvfgfsk9n9
2gkbqpqn31
sevenfivefgcvtpsxjprfh8nine
lvfxml6992
5threexmjjgkv8985zzjqdbtwo
eight3bkqdnbmrtb5
7sevenoneznpx
9hjg4eightrcsvdkbmldjclfqfm
9four2one4
koneightonecdfcrjkqtcsevennlvr8hbrc
seven3lthbv8
ninexzznsix5nine
ninentdd6qvkclninefivenine
four11pgmxz7pnjfiveeight
threeninefourpmtmlgllftnvxzn5twonine
four46eight9sevenfive3three
blztvfkqggcbshlzxppxgrxsr4three3
one9pvlnv
8l37
brjvlvjrhbfourxshh1vlclvrz2
5flqnkntnxkzhcftmzb9rlfzxlg
sixfourseven77onekone
6xfmdgjfeight4qzrncxdpmb4four
four28
seven74zrtpftdldc7fqlseven5
6gxjzskpkfvmmd
zkoneight99jrrmgsfpsixfiveone
9jbbdtdxjsm8szxblgjppx4tpnvqvtlrj
rsrlrcb32ngsixfourfivenine
6seven9one7threegrfqpncjthree
onebjvpzzqhvlhg1pxkvmgqvxsrglb8fourptbjs
1fch
njhs4p86n22
two1one6four1six
fourphxjkjtwo8eight
3svqrzd1fourcmlcknhvninetddpbcmbtgqpcjms
vkndzm684sixjlgkcvz4
25five
pkzt2fivetwo2zffkjqrhgfive4jpsj
sevenmpfcthreembjgbfpkdzqlr4
49four1mpldzb3
hdzddkxf1cxftflb73
zrmhdlhk2v7
25zcd
11eighteight5qcqzpvvk
38nine43
8lstwo
12sixnfbrgbhdpn1three3
8eightcd82rzkzlvthqnvhjvgfour
24jcbjplcnqbcrxs25
3fivebfkgpkglfchbmbfps
pgmvbnhskgzdmz2sixeightcjq5
cdpmrlj1one
6hxqoneightjjv
34gxbjzrtg5
pvddskbslqnrfngmcjgsdthree3
6nkjjlknp9
hflkjhgjmeightc5n1
dmkeight1223
five9three8
eight7mqfsjplfprrfpkzctgtdvrmxphm
phbs2fck4sixfourqvqbr
2bnvktn
sfvkgzone83
2threetqnthdnq32ninetwo9
sixhdkvdcmp5three23j
8615four
twosix54vrrbvzszk9
53671
threevklcphgkjsnine4eight4fmtffknglthree
sixninethreeh4dj
oneone1pmdthreesrfsssbkt8
7foureighteight3
srpvkzrqfive378
6twodnnrvfjrjv
4kpxlslqfbktwo
bvzpgrc9twotqsvdztwoseven8
xkjdltjgzbjhxkjvtwo6
five27sixsevensixtwo9nine
jk468qgkr
fourtwo89
fourthreethreedtnzbmlfhmgjr5nineseven7
44seven
two3psfive122jps
sevensevenghzmpdvrffive9nine3eight
497ftdf9five
gklfive6rnvpnvvkqheight
2hvdfiveprrdqspsix6
5m8nine
bnctbninexsixonesix8five8
xkqqlmfmrveightsix4nine93nine
6k
3foursevensix6cksix5six
7twoone
three73lfddzhd1fkxmjdzsix7twonex
sthjlrjrhd741prcsqh3rmllvjmtvgfour
4dvffpjkn
tpbttcslvz7twoneq
fourtwoninej5snfxnqzthree
n7cmsfsqd
eightssrzkxj5
239
three8five7xxthreebqrbx
zrlchvsevenfournine4ktvskhjgh3h
7fourfdjsnhdbgqjvnltzj6three
mrjsndmzkz7rszqnbhxt3fvsix1
neight85eightggtnxtgljsevenfivekz
6vknslh4onetwonrlzm
37zlrksix1skbsdkpjf6twonejtx
three2843five6
txgdvvdg959
zfjrs2zvtbqctcdqgrpfmqsjbdone5
six9eightninetwo72sxxnzvblthree
2jdpslvbnpqjpglczkmzggkfkdkx8hgpxtcz
threefive5eight5mnbllfpcsp29vlnbrntt
gfxndggbs16twozpcsckzqcj3sthsgq6
gcjjvqkvzdbcsnmqqhnzzqvj4
flcpl3btfmbbpnkjvnlmcthreetwo1eightwops
bdmeight67tvkfh2
three645qcv1zbbheight
3ninethvbxxppxgqcqrclptxczgrcneighttwofivebrqxl
25eight41
six22
pcp5
dtmgxkdqsixdhmsbj821
eightprbxpj5oneightcxj
qvrn3jbhlxjsdq
oneclvhjhr5
9nkmqpjjxxhvtpndls9
215ltwo75
threeeight16nine2mzhxnine
ktfxkmdvzprhkpdhvxhzsc68
fivesix8five
4z1eighttwofive
73twotwo4
gmjlpchdzfthreesix1vljxdqsrlxmmqs2
twofxh3
threeqzcglsdcfm4four
three1eight8
seven7dtqhr7
4xbjlxlptj8hzfjnz
37jzgxbjcggone
81fouronenine489four
eight3fiveoneseven135
stzts59zqdvrdcqrc
sixnklrjbeightn2six
bqdtntwonine1eightttzlzvzfn54xmj
sjk6
qmrbnhczj624394sevenseven
kkjdcjhfh93eight
3ptmslnconethree
cfjgdffcgvldsnvkbjqrxhxcl7fjlxdrlrrthreeseventwo
187oneseven6
fiveprnppdg2tjfbfmlvhpmkggjc
9onetwo4
fivesvjxkzzm59vtsevenhzxtkggdhr8hvjtjvv
kpxkbbxseveneight89sevenrbhqqpk
drhkpssxqvvnssq59four4
oneeight17
gteightwoone268four
eight7fourbjnlzfiveczlzppnxck
jdqpxsevenone2eight
pvh5six4hddrhfzpxfmtwo
5six56nineone2
v96k9115three
phqhhthree5ksqhfjlbfg
gpsskbfhhllnxfvjspkjndtlfour5
three5zqbnsrdthree8
b4mkfpkltlfdfive4mdqxjnb8tdpnpf3
qfivejjggrpktxponetrjzceightseven9xhdf
34cpfxc
jl9
84four
fourthree9three48
7threevgvtj2five4pbq
7kjkjc
seventfsvjbsh5smmdd3frthree
smgmzqzn5
gmtd5kvglxsixeight8twothvkprlbc5
twofour62
nine5k
hbxnpb4four4h
6sixgnmnjv4fourone1
9eightseventhree
9jpvccsvhqpnhsl8
nvcnninefour9
v237ppqbhb
78six
four165oneightxcm
9pjcsfbrghnineqzth4smx
46fpfptrq1mbqmbnktqeight
stbxvlcqz5krd1threethreeonefour
48six5seven
6bmltlrvrgpcfhjhmfiveqzfxptjtwo4zvsqqxgbrdlzsfmtzdd
jeightwo47three86twoseven
njxzmthree8fivevvchvjqdvn3foursix
83t34
oneznzqptpxbrtqxstkmz3kmtstds
1sixsfrnqd
brs7fgkbhntv5s5
twohs7m
gfljsixseventmgdvhqthree3threefive
f1twonekdh
oneone735eightnine
pfjvfspsseven9qgfrnmckxzsix94
7nineninelrcqqcgcnmmqf3
1vjdhjtrfourfive2nine278
oneone7
threelcxlqrzhdghp4zkjfivepjj
hcshggsmzpdmkvtdvdgqtfxlt8
1vdjtvpfhkhfive1
2eightstnj6three
93threeeightntjblpljbv
onermlmtncmj6pxkmhmqchzvzf
4eightfnjzfzhvg
hmjvmtwovnl8nine84
threets6
six9twofivexgz63
692five
fkgblptntvhvmlv8threethreenine
seven86
nine1n7
97two
4threebqqnmvmqleightthreefiveszppmbhxrxvpxz
6bxbmbdkxqmzeightlrqdqvrkr5threethree
llqcbpeight1vpjninekpdvzg
6qgtdqvtvkcbcskfqtq
6threegrrmxxxqkflltr3
312
hqmhsxpmkxtwosix3
rpzqtmzgdfdxcgsix1six63lxjpbxfq8
zzkbtkghmmqfourrtsixxxfjnvvccmpsd5six
13two4bxdcqzrkqtxm1mplvqxcfhcjsc
threesixqj8two
9seventhree
dtnine5twoseven6zxd
lsqbvgjnznineone7lxtvmkmflrfcqdjmjtwo
zpkjlfp6onevxtdtdzmcqjprfive
gxplqqsz412sevenninejs1
psvjsvvnrv796
69tfxkbkchvlhkjbrmone
fivehnrvtb6
eightzdlqrbzxteightptlgmcmvtwothreergcddqxf2twonepxh
343sevenxsffneightdvft
fivednmrpmvv8fiveninesevenzbggk
seventwot3hpfrzbhxlhfivetwo7zvmpmq
xnsxz8fivezhzdcbzsvp
sfzch8twoned
5rstpx
12six
gqznine5gpg
6zrmsp825seven
2493twothree
9fivemksdnmgbvx
6dvdpdpkmqpxvfive28six
dzvnrdksixonetwoonetr4
7fncndxbqj6onetwosixsixthree
7cnprcdgk2three2dvtccqnskvzfsevenxdrnqf
cxfcdvbsjqjbnxddlggjfourtnfzvtgx4twoddkkpdd
pxjgqrmdg5mpcgcdmfeight825jxxqcnfive
1tkhgtzzfrbdvnbft
hgxfive14ddcfhshfd4
onenzlhvtdgkjmjgldmddhngdv9onebkt
foureight8rzxkktk9eighteight1
478nshqhnhjrmlqbmp
ddjzzxgj75zktccgqrltfivethree
eightmcnmt5jnmnqhqdfive
rcdxshk1seven5
rdcmbznk79
2gtbskjxpmmvdclgmfjrc48one
8mqgnfive7chknsixrfourseven
4hgdxjgbn1sixseven7twosixseven
mkjslkltjd59
fourfoursixfkjrcfsfivenrtzv4eight
five8threeseven4
81821tcmfourddhmzvzfive
onesix943fivejld
5eighteight3cmvvpqmdq72vrvb
two86
nqjrpqvgqr7rjjjxglqzrmt
5kqjjvzxt
6fourvstjrlnvone3ndphzphkrnsqmkmsthreep
six8four7
two1dntwo
eightgmcgrnptrcvztbdp4three
xxbpnnztr5eightpnqeight
three9pfpnjx6rkphpjeight3five
sevendzstsjl3krspscb1
fsevenfqtxxhjzvnineninesixeight8
kznjhnxbnk7qbxjrztltv
rkeightwo6zfpvrfgqr7qxbkkg43lrjqtzjrprqttxmbrzg
mcfive77vgzxonehglbj
9319
4mzds
sevensixthree5sixdvzxkndhvjfive
9mjhfkeight88v
seven5zhdthbmrkdpdxfcp8njtqvpnjj14
8ninetwo
95tnjldjqcrzdxlm3
fourgffour8
nine821qbv6five
jbtfg83two
twozcpjrcnplnz5bdtgpdctb87lzlvqhtrjj
tr4
5foureightfourfcs
1cqjts1jgzkfm
66threetwo
jsdpkfnineeightzpjdmrvxkbhdntj9
pgcqrsix6mqrr8threeqxgkftbmzninevndn
kmvqsqhbrcnbqqgninet6
bxtstqzpqfzqnhjfb8htszvgqhpnggvqt
ninecpqpffivehg8
kxmstxkffourqmx41
scfourlkfbrjvbtwo5
qr88fivenine1lfvksgrtqseven
four6foursix
twoqmnxrjql5fourpdlstnnsfkdjgt9
4zsvbsjqv97bpxvncr
xhqlhsbqjhvdqqonesevenfive3qvrtbkhhlfbzsj9
sixeight4six5szgzcjhpj1
7vpjq9hjtrjgone
f92eight
eighteightnc8134
22onespjpxlttsqsix8eighthxdfvsdx
5ninefourgrrmxsxjfxk
threeninepmvsv763xlxjp
oneone7lzzhjqqrg
9fourtqqmhrpmkxhrvlnjvvhsevenseven
threethreetvjpnoneseven8
two46onetwoqbvntlxbrftpjf9
8one5nine1three8
three69sgdkstpqbqdz
q8rzcl
onepjmchxtlqnmrcrvm6
61fivegjjsevenqgdkq
49fxhdzfntmk6tb8dpdkknzsb
bbseven1xvqmlrhx
2j
seventwozjqszlhzxlpgphnkz2foursixfour
1sixhgvhrbonetwom
fournine8gvmrpgdxvcbdspzdcqt87bdzvxbf5
3sevennine2fzpt
threedcdlq9kcjhtmtz
86five
tpfqhqs3977
8four1cgmm12shfl58
xlkrrkpkqjtslblqfnxp7two16jzpmpkrfvdzh
243
three9hkgnmrh6lqrsx5
fivexsczpmltrmcgrvfc58
mxngrsh2sevensixthreelgrmg
ljqmflvone76
ksponeightthree2ninenine
tdsdmjznr5nine7fourtwojgjsdfsevenone
eightfive1fourseven3gsqhtv
1xfdmqtmgkmjkthree
rztwonelztpgkxzzcbn1eightttssdpone71
zthhsgvmhqsgvdponetwo9p
293rmjjjpmjchjnbdcssfrneightvdzrkbhdln
7nfkdntfourthreefzrfxmxgqone5
tkgrnhbflp7zltmbdoneeightwoh
7bspgfklffgsix4
eighthcmlrpbhjjmvbjrleightgd94
sixbgfjzgzbxsb4qsixthreehbbn
ninerhzjpjdfnsevenonenine8
38kqzjxqmmm
jngngvc412
4twosponesixdpj15five
41four2oneonekr2
cpsixjnlhkthree353seven
bknflgv1sixfivefive5
3qkhnsjqkcjmlg7gl4jthreethree
dhlngstrvbxjjll2979kjsttsfgjkc
4kkq7rqlxcldqqtwo
jfh74
lkfpcdghgq8gpgldrsnzkzzzzskrvcvsjthree
57fivefive3cxqj2
ztwone7vcd37122
4qrrhhlxgpr4
twokzfjg2sevenlnvlpzxknznpsc1
4jjbcdbfm8six6four
5eightninesixvzvf98two
8flntwomkktkpvsone78sixone
sevenzltjhkptjfjbrppm85eight
fivemfrmnqptthreepninepd5
6one9
vkqxgzmbm2b4pjqjddsbjnjcqqvm9
5443nfkv
leightwo5
8fiveeightonetwovgvhzgzfjh16eightwohlk
fivexnbhkzjfg1
chkxvgrgb1sqxsnhngnrtqsnqgjkd
96twoone5553dv
36pfltskrbcmlmnspn
4jbbrh95249
xfzspqssdfourhnmtzfive2pfzczh
onexzbzhddkqgfr2
sevenninerrlveight5nine3
eight7qvgkbk238fiveeight
sevenseveneightgtxtxkjsgdgklzzxxc3two8
tvfjhvtclm75skqdxsskqhrjkbg
45
sevenfive82
425zkhjhmk
onexdchhtxmhsevenbczrslrppneightonenbnhfmbsvdcnzjx1
zfkscdbmtwoeightrksdmgx4
lnseightnine9eight
fourhmbhlcpht53ngkbzjmfivesixg
threestrhbj9sixggczcg
9twoskgrps8
3four6xdqczgtzlzf
ldfn2
qlzjsnbzfourfdq476
tpkczdh5hdbxvvmmt3sixsix
tdpxzld5
lnveightwohdkgcvvrjs38
zz8eightstvmhvrh7hftdhkrjcneight1
9kdbcpqtx15
rbqgdbvrstgninefive4bqq2six
nine3psmkzkgnjbndrcninesevenzvcr6
eight4jfrqcbqfninedxmdtjgqgtrg6four5
3mmnineninenjjpmfivetwo
cv4znxcjthreeqqtdqmzxfknnp
8pjkm
ptwonethreegrgvseven7
onethree1bgjsix5sevengpts
96xlmmthreeeightcbdnrstvpncmr
rtc94tcninefive
onespqnnptpdbrgqsqrldstl1
qkeightwotwocjcngknkztwo7
gzjhzlf4fdglcrzckbrlkmg
3eighttwopninefour
2seven1c
17nine447qkmfour2
six6xfgqddnfpsc
mkbgbkvzdpzxfmrhdcjklxfoureightzzpn3eight
95ninevhctbgznbzz871sixoneightr
31onekmseveneight2four
mlxqgmvj2six6
2sixgvsbmrhtwofour
eight221three99two3
88xrrbjdlzrfour8plv
xdglmrpxbz5xpjxzpmvrgsixthreeseven7threebtqfkqp
lbd2onethree
seventwoseven7threesixbpld
1pstwofour8eight3dsdfrseven
gnvzm19htsbvcsfmlrmbgtstzmm3twoqzffkrrq
94nvrbbj
one71rsfbpnnbkrklmxqfive
4six1
eight48chsrmsix
vqxrnmsix98hlzdgvd3sevenninekng
12threehscqzvzcbgfive6three2zhtthr
6sixkzrnv
5gmnhhzkfmp
four35seven7onenvdsevenftnpbcj
6zxrhcxxkppkn2
dgshxchmhgtgjk281seven
lbdsmfvdsfzlp6dfpgd3
three2dpsdhfld95eightwoht
jbktdklsqkgnhnfmseven1lhdsbjksixtwo
3fdqfour
5jpljkkpmdsix
qfeightwo9threethree
hpqdx4911tzfcxlrtccqf9one
q79zspgmjpdzs63
344zk5xbthreezgbffcb
nlzmjfqxmneightxqjdnjvr21
8cxtrkpvzj21xfgbdgcvgrztwo
three18444
sntrptktwo2one1five
qxjrgfcnpcjtnfjljqnq1onehzfcqlnine
kflgzv58dbzbjjdvclgtseven
twofive4eightwozz
eightfive365
7nqnksvphhnine88
t8eight
bjd6five
khbrbtsx5jqxmbsqtf5nine3sevenskhfg6
seventlkmfhqkgxkbhqr6ncjztnfive
152one
three5cbpqkzb4eightseventgmqzflsfksix
seven32threegfddgtf
two5twofivexzkkvcqs3
tvvdgrnqlmkfour1zrcznqkhseveneight1q
92threesix89
3zbdlttpbh7fivepgxmrvbzlnfgmbkzknndfqk5
threefnhxtdbl1jtxeightwol
5vgthldgfmgdjphvcgh53dshmdkc
3nineeightwokh
57eighttddbcdsjdss
tkcgn86xfgbmzt7rksvnchnrh
five3dn5
three2eighteight15nine
49bn1zvbm57
351six6xfzfjvpz
5one5zchddj4dkksn
xfhtwonesevenfivethreepqzmrzrrfourthree5
ninetzmcgp47four
sevenbcfbpnrvkkscrjtpctdtb69bvvnvlgsmjltlvs
6threev
ninerlsbznvfn9
fourbm2
sdxd22
n7
7sixcjdsxfourfmvzrbvlnine5
threeqtbhgznine7one
ftmkmxkd9fvvlg353rp51
9zjhmpnjv5jvndz
58three59nineonesix
rmjvhrjjmkqsn6gqthreeonefivemxqhrzvffone
xsslv7gpgkbzdmr434four
pnzxp4nbtsjqctkvqncxzxzj
eightseven52five4ninekntfjrdt
4sixldsmv
pknxkqgdpnc7fivedbvhkn
qkpjhjlxone4sixpfkvhlmxmd3
four3ninerkrcvgcmbb2qm
fivenine6six1eight
69sixnine
bvjx5lg5vgrqq
21ninegnhdkcxhzkcfdksvsmdthree
zjrnmhclxhrkjpffhxkthnvj83jnshbqvx
bzfphcg9fourthreegkchdvrgsx
2ninebvgdcfxtktqjxjqvxfgjdqfhv5threegqtsfhtfxg
6rqskvckjzq2qzrnbxjmlthreeeight6hrs
sixthree6lxcrsevenseven69twonegs
2dcvcqcbpshsixone3
drkdbmv4zbjbznsqtj
eightbqfhnmvqsoneninezbrzcqkz4ftv
1eightcrcjcbdthreebscfpvznqfrj6
`
var sum = 0

func main() {
	var lines = strings.Split(input, "\n")
	var sum = 0
	lines = lines[1 : len(lines)-1]
	for _, line := range lines {
		numberLastIndex := strings.LastIndexFunc(line, unicode.IsNumber)
		numberFirstIndex := strings.IndexFunc(line, unicode.IsNumber)
		firstAndLastStr := fmt.Sprintf("%s", line[numberFirstIndex:numberFirstIndex+1]+line[numberLastIndex:numberLastIndex+1])
		firstAndLastInt, _ := strconv.Atoi(firstAndLastStr)
		sum += firstAndLastInt
	}
	println(sum)
}
