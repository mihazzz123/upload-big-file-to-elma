# currency
`import "git.elewise.com/elma365/common/pkg/types/complextypes/money/currency"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Currency](#Currency)
  * [func Parse(n string) (Currency, error)](#Parse)
  * [func (c Currency) MarshalJSON() ([]byte, error)](#Currency.MarshalJSON)
  * [func (c *Currency) UnmarshalJSON(data []byte) error](#Currency.UnmarshalJSON)


#### <a name="pkg-files">Package files</a>
[currency.go](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/currency/currency.go)



## <a name="pkg-variables">Variables</a>
``` go
var CurrencyName = map[Currency]string{
    Currency{"Afghani", "AFN", "AFGHANISTAN", "971", 2, "\u060b"}:                                                                        "AFN",
    Currency{"Euro", "EUR", "ÅLAND ISLANDS", "978", 2, "\u20ac"}:                                                                         "EUR",
    Currency{"Lek", "ALL", "ALBANIA", "008", 2, "\u004c\u0065\u006b"}:                                                                    "ALL",
    Currency{"Algerian Dinar", "DZD", "ALGERIA", "012", 2, ""}:                                                                           "DZD",
    Currency{"US Dollar", "USD", "AMERICAN SAMOA", "840", 2, "\u0024"}:                                                                   "USD",
    Currency{"Kwanza", "AOA", "ANGOLA", "973", 2, ""}:                                                                                    "AOA",
    Currency{"East Caribbean Dollar", "XCD", "ANGUILLA", "951", 2, "\u0024"}:                                                             "XCD",
    Currency{"Argentine Peso", "ARS", "ARGENTINA", "032", 2, "\u0024"}:                                                                   "ARS",
    Currency{"Armenian Dram", "AMD", "ARMENIA", "051", 2, ""}:                                                                            "AMD",
    Currency{"Aruban Florin", "AWG", "ARUBA", "533", 2, "\u0192"}:                                                                        "AWG",
    Currency{"Australian Dollar", "AUD", "AUSTRALIA", "036", 2, "\u0024"}:                                                                "AUD",
    Currency{"Azerbaijan Manat", "AZN", "AZERBAIJAN", "944", 2, "\u20bc"}:                                                                "AZN",
    Currency{"Bahamian Dollar", "BSD", "BAHAMAS (THE)", "044", 2, "\u0024"}:                                                              "BSD",
    Currency{"Bahraini Dinar", "BHD", "BAHRAIN", "048", 3, ""}:                                                                           "BHD",
    Currency{"Taka", "BDT", "BANGLADESH", "050", 2, ""}:                                                                                  "BDT",
    Currency{"Barbados Dollar", "BBD", "BARBADOS", "052", 2, "\u0024"}:                                                                   "BBD",
    Currency{"Belarusian Ruble", "BYN", "BELARUS", "933", 2, "\u0042\u0072"}:                                                             "BYN",
    Currency{"Belize Dollar", "BZD", "BELIZE", "084", 2, "\u0042\u005a\u0024"}:                                                           "BZD",
    Currency{"CFA Franc BCEAO", "XOF", "BENIN", "952", 0, ""}:                                                                            "XOF",
    Currency{"Bermudian Dollar", "BMD", "BERMUDA", "060", 2, "\u0024"}:                                                                   "BMD",
    Currency{"Indian Rupee", "INR", "BHUTAN", "356", 2, ""}:                                                                              "INR",
    Currency{"Ngultrum", "BTN", "BHUTAN", "064", 2, ""}:                                                                                  "BTN",
    Currency{"Boliviano", "BOB", "BOLIVIA (PLURINATIONAL STATE OF)", "068", 2, "\u0024\u0062"}:                                           "BOB",
    Currency{"Mvdol", "BOV", "BOLIVIA (PLURINATIONAL STATE OF)", "984", 2, ""}:                                                           "BOV",
    Currency{"Convertible Mark", "BAM", "BOSNIA AND HERZEGOVINA", "977", 2, "\u004b\u004d"}:                                              "BAM",
    Currency{"Pula", "BWP", "BOTSWANA", "072", 2, "\u0050"}:                                                                              "BWP",
    Currency{"Norwegian Krone", "NOK", "BOUVET ISLAND", "578", 2, "\u006b\u0072"}:                                                        "NOK",
    Currency{"Brazilian Real", "BRL", "BRAZIL", "986", 2, "\u0052\u0024"}:                                                                "BRL",
    Currency{"Brunei Dollar", "BND", "BRUNEI DARUSSALAM", "096", 2, "\u0024"}:                                                            "BND",
    Currency{"Bulgarian Lev", "BGN", "BULGARIA", "975", 2, "\u043b\u0432"}:                                                               "BGN",
    Currency{"Burundi Franc", "BIF", "BURUNDI", "108", 0, ""}:                                                                            "BIF",
    Currency{"Cabo Verde Escudo", "CVE", "CABO VERDE", "132", 2, ""}:                                                                     "CVE",
    Currency{"Riel", "KHR", "CAMBODIA", "116", 2, "\u17db"}:                                                                              "KHR",
    Currency{"CFA Franc BEAC", "XAF", "CAMEROON", "950", 0, ""}:                                                                          "XAF",
    Currency{"Canadian Dollar", "CAD", "CANADA", "124", 2, "\u0024"}:                                                                     "CAD",
    Currency{"Cayman Islands Dollar", "KYD", "CAYMAN ISLANDS (THE)", "136", 2, "\u0024"}:                                                 "KYD",
    Currency{"Chilean Peso", "CLP", "CHILE", "152", 0, "\u0024"}:                                                                         "CLP",
    Currency{"Unidad de Fomento", "CLF", "CHILE", "990", 4, ""}:                                                                          "CLF",
    Currency{"Yuan Renminbi", "CNY", "CHINA", "156", 2, "\u00a5"}:                                                                        "CNY",
    Currency{"Colombian Peso", "COP", "COLOMBIA", "170", 2, "\u0024"}:                                                                    "COP",
    Currency{"Unidad de Valor Real", "COU", "COLOMBIA", "970", 2, ""}:                                                                    "COU",
    Currency{"Comorian Franc ", "KMF", "COMOROS (THE)", "174", 0, ""}:                                                                    "KMF",
    Currency{"Congolese Franc", "CDF", "CONGO (THE DEMOCRATIC REPUBLIC OF THE)", "976", 2, ""}:                                           "CDF",
    Currency{"New Zealand Dollar", "NZD", "COOK ISLANDS (THE)", "554", 2, "\u0024"}:                                                      "NZD",
    Currency{"Costa Rican Colon", "CRC", "COSTA RICA", "188", 2, "\u20a1"}:                                                               "CRC",
    Currency{"Kuna", "HRK", "CROATIA", "191", 2, "\u006b\u006e"}:                                                                         "HRK",
    Currency{"Cuban Peso", "CUP", "CUBA", "192", 2, "\u20b1"}:                                                                            "CUP",
    Currency{"Peso Convertible", "CUC", "CUBA", "931", 2, ""}:                                                                            "CUC",
    Currency{"Netherlands Antillean Guilder", "ANG", "CURAÇAO", "532", 2, "\u0192"}:                                                      "ANG",
    Currency{"Czech Koruna", "CZK", "CZECHIA", "203", 2, "\u004b\u010d"}:                                                                 "CZK",
    Currency{"Danish Krone", "DKK", "DENMARK", "208", 2, "\u006b\u0072"}:                                                                 "DKK",
    Currency{"Djibouti Franc", "DJF", "DJIBOUTI", "262", 0, ""}:                                                                          "DJF",
    Currency{"Dominican Peso", "DOP", "DOMINICAN REPUBLIC (THE)", "214", 2, "\u0052\u0044\u0024"}:                                        "DOP",
    Currency{"Egyptian Pound", "EGP", "EGYPT", "818", 2, "\u00a3"}:                                                                       "EGP",
    Currency{"El Salvador Colon", "SVC", "EL SALVADOR", "222", 2, "\u0024"}:                                                              "SVC",
    Currency{"Nakfa", "ERN", "ERITREA", "232", 2, ""}:                                                                                    "ERN",
    Currency{"Ethiopian Birr", "ETB", "ETHIOPIA", "230", 2, ""}:                                                                          "ETB",
    Currency{"Falkland Islands Pound", "FKP", "FALKLAND ISLANDS (THE) [MALVINAS]", "238", 2, "\u00a3"}:                                   "FKP",
    Currency{"Fiji Dollar", "FJD", "FIJI", "242", 2, "\u0024"}:                                                                           "FJD",
    Currency{"CFP Franc", "XPF", "FRENCH POLYNESIA", "953", 0, ""}:                                                                       "XPF",
    Currency{"Dalasi", "GMD", "GAMBIA (THE)", "270", 2, ""}:                                                                              "GMD",
    Currency{"Lari", "GEL", "GEORGIA", "981", 2, ""}:                                                                                     "GEL",
    Currency{"Ghana Cedi", "GHS", "GHANA", "936", 2, "\u00a2"}:                                                                           "GHS",
    Currency{"Gibraltar Pound", "GIP", "GIBRALTAR", "292", 2, "\u00a3"}:                                                                  "GIP",
    Currency{"Quetzal", "GTQ", "GUATEMALA", "320", 2, "\u0051"}:                                                                          "GTQ",
    Currency{"Pound Sterling", "GBP", "GUERNSEY", "826", 2, "\u00a3"}:                                                                    "GBP",
    Currency{"Guinean Franc", "GNF", "GUINEA", "324", 0, ""}:                                                                             "GNF",
    Currency{"Guyana Dollar", "GYD", "GUYANA", "328", 2, "\u0024"}:                                                                       "GYD",
    Currency{"Gourde", "HTG", "HAITI", "332", 2, ""}:                                                                                     "HTG",
    Currency{"Lempira", "HNL", "HONDURAS", "340", 2, "\u004c"}:                                                                           "HNL",
    Currency{"Hong Kong Dollar", "HKD", "HONG KONG", "344", 2, "\u0024"}:                                                                 "HKD",
    Currency{"Forint", "HUF", "HUNGARY", "348", 2, "\u0046\u0074"}:                                                                       "HUF",
    Currency{"Iceland Krona", "ISK", "ICELAND", "352", 0, "\u006b\u0072"}:                                                                "ISK",
    Currency{"Rupiah", "IDR", "INDONESIA", "360", 2, "\u0052\u0070"}:                                                                     "IDR",
    Currency{"SDR (Special Drawing Right)", "XDR", "INTERNATIONAL MONETARY FUND (IMF) ", "960", 0, ""}:                                   "XDR",
    Currency{"Iranian Rial", "IRR", "IRAN (ISLAMIC REPUBLIC OF)", "364", 2, "\ufdfc"}:                                                    "IRR",
    Currency{"Iraqi Dinar", "IQD", "IRAQ", "368", 3, ""}:                                                                                 "IQD",
    Currency{"New Israeli Sheqel", "ILS", "ISRAEL", "376", 2, "\u20aa"}:                                                                  "ILS",
    Currency{"Jamaican Dollar", "JMD", "JAMAICA", "388", 2, "\u004a\u0024"}:                                                              "JMD",
    Currency{"Yen", "JPY", "JAPAN", "392", 0, "\u00a5"}:                                                                                  "JPY",
    Currency{"Jordanian Dinar", "JOD", "JORDAN", "400", 3, ""}:                                                                           "JOD",
    Currency{"Tenge", "KZT", "KAZAKHSTAN", "398", 2, "\u043b\u0432"}:                                                                     "KZT",
    Currency{"Kenyan Shilling", "KES", "KENYA", "404", 2, ""}:                                                                            "KES",
    Currency{"North Korean Won", "KPW", "KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)", "408", 2, "\u20a9"}:                               "KPW",
    Currency{"Won", "KRW", "KOREA (THE REPUBLIC OF)", "410", 0, "\u20a9"}:                                                                "KRW",
    Currency{"Kuwaiti Dinar", "KWD", "KUWAIT", "414", 3, ""}:                                                                             "KWD",
    Currency{"Som", "KGS", "KYRGYZSTAN", "417", 2, "\u043b\u0432"}:                                                                       "KGS",
    Currency{"Lao Kip", "LAK", "LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)", "418", 2, "\u20ad"}:                                             "LAK",
    Currency{"Lebanese Pound", "LBP", "LEBANON", "422", 2, "\u00a3"}:                                                                     "LBP",
    Currency{"Loti", "LSL", "LESOTHO", "426", 2, ""}:                                                                                     "LSL",
    Currency{"Rand", "ZAR", "LESOTHO", "710", 2, "\u0052"}:                                                                               "ZAR",
    Currency{"Liberian Dollar", "LRD", "LIBERIA", "430", 2, "\u0024"}:                                                                    "LRD",
    Currency{"Libyan Dinar", "LYD", "LIBYA", "434", 3, ""}:                                                                               "LYD",
    Currency{"Swiss Franc", "CHF", "LIECHTENSTEIN", "756", 2, "\u0043\u0048\u0046"}:                                                      "CHF",
    Currency{"Pataca", "MOP", "MACAO", "446", 2, ""}:                                                                                     "MOP",
    Currency{"Denar", "MKD", "MACEDONIA (THE FORMER YUGOSLAV REPUBLIC OF)", "807", 2, "\u0434\u0435\u043d"}:                              "MKD",
    Currency{"Malagasy Ariary", "MGA", "MADAGASCAR", "969", 2, ""}:                                                                       "MGA",
    Currency{"Malawi Kwacha", "MWK", "MALAWI", "454", 2, ""}:                                                                             "MWK",
    Currency{"Malaysian Ringgit", "MYR", "MALAYSIA", "458", 2, "\u0052\u004d"}:                                                           "MYR",
    Currency{"Rufiyaa", "MVR", "MALDIVES", "462", 2, ""}:                                                                                 "MVR",
    Currency{"Ouguiya", "MRU", "MAURITANIA", "929", 2, ""}:                                                                               "MRU",
    Currency{"Mauritius Rupee", "MUR", "MAURITIUS", "480", 2, "\u20a8"}:                                                                  "MUR",
    Currency{"ADB Unit of Account", "XUA", "MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP", "965", 0, ""}:                       "XUA",
    Currency{"Mexican Peso", "MXN", "MEXICO", "484", 2, "\u0024"}:                                                                        "MXN",
    Currency{"Mexican Unidad de Inversion (UDI)", "MXV", "MEXICO", "979", 2, ""}:                                                         "MXV",
    Currency{"Moldovan Leu", "MDL", "MOLDOVA (THE REPUBLIC OF)", "498", 2, ""}:                                                           "MDL",
    Currency{"Tugrik", "MNT", "MONGOLIA", "496", 2, "\u20ae"}:                                                                            "MNT",
    Currency{"Moroccan Dirham", "MAD", "MOROCCO", "504", 2, ""}:                                                                          "MAD",
    Currency{"Mozambique Metical", "MZN", "MOZAMBIQUE", "943", 2, "\u004d\u0054"}:                                                        "MZN",
    Currency{"Kyat", "MMK", "MYANMAR", "104", 2, ""}:                                                                                     "MMK",
    Currency{"Namibia Dollar", "NAD", "NAMIBIA", "516", 2, "\u0024"}:                                                                     "NAD",
    Currency{"Nepalese Rupee", "NPR", "NEPAL", "524", 2, "\u20a8"}:                                                                       "NPR",
    Currency{"Cordoba Oro", "NIO", "NICARAGUA", "558", 2, "\u0043\u0024"}:                                                                "NIO",
    Currency{"Naira", "NGN", "NIGERIA", "566", 2, "\u20a6"}:                                                                              "NGN",
    Currency{"Rial Omani", "OMR", "OMAN", "512", 3, "\ufdfc"}:                                                                            "OMR",
    Currency{"Pakistan Rupee", "PKR", "PAKISTAN", "586", 2, "\u20a8"}:                                                                    "PKR",
    Currency{"Balboa", "PAB", "PANAMA", "590", 2, "\u0042\u002f\u002e"}:                                                                  "PAB",
    Currency{"Kina", "PGK", "PAPUA NEW GUINEA", "598", 2, ""}:                                                                            "PGK",
    Currency{"Guarani", "PYG", "PARAGUAY", "600", 0, "\u0047\u0073"}:                                                                     "PYG",
    Currency{"Sol", "PEN", "PERU", "604", 2, "\u0053\u002f\u002e"}:                                                                       "PEN",
    Currency{"Philippine Peso", "PHP", "PHILIPPINES (THE)", "608", 2, "\u20b1"}:                                                          "PHP",
    Currency{"Zloty", "PLN", "POLAND", "985", 2, "\u007a\u0142"}:                                                                         "PLN",
    Currency{"Qatari Rial", "QAR", "QATAR", "634", 2, "\ufdfc"}:                                                                          "QAR",
    Currency{"Romanian Leu", "RON", "ROMANIA", "946", 2, "\u006c\u0065\u0069"}:                                                           "RON",
    Currency{"Russian Ruble", "RUB", "RUSSIAN FEDERATION (THE)", "643", 2, "\u20bd"}:                                                     "RUB",
    Currency{"Rwanda Franc", "RWF", "RWANDA", "646", 0, ""}:                                                                              "RWF",
    Currency{"Saint Helena Pound", "SHP", "SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA", "654", 2, "\u00a3"}:                            "SHP",
    Currency{"Tala", "WST", "SAMOA", "882", 2, ""}:                                                                                       "WST",
    Currency{"Dobra", "STN", "SAO TOME AND PRINCIPE", "930", 2, ""}:                                                                      "STN",
    Currency{"Saudi Riyal", "SAR", "SAUDI ARABIA", "682", 2, "\ufdfc"}:                                                                   "SAR",
    Currency{"Serbian Dinar", "RSD", "SERBIA", "941", 2, "\u0414\u0438\u043d\u002e"}:                                                     "RSD",
    Currency{"Seychelles Rupee", "SCR", "SEYCHELLES", "690", 2, "\u20a8"}:                                                                "SCR",
    Currency{"Leone", "SLL", "SIERRA LEONE", "694", 2, ""}:                                                                               "SLL",
    Currency{"Singapore Dollar", "SGD", "SINGAPORE", "702", 2, "\u0024"}:                                                                 "SGD",
    Currency{"Sucre", "XSU", "SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS \"SUCRE\"", "994", 0, ""}:                               "XSU",
    Currency{"Solomon Islands Dollar", "SBD", "SOLOMON ISLANDS", "090", 2, "\u0024"}:                                                     "SBD",
    Currency{"Somali Shilling", "SOS", "SOMALIA", "706", 2, "\u0053"}:                                                                    "SOS",
    Currency{"South Sudanese Pound", "SSP", "SOUTH SUDAN", "728", 2, ""}:                                                                 "SSP",
    Currency{"Sri Lanka Rupee", "LKR", "SRI LANKA", "144", 2, "\u20a8"}:                                                                  "LKR",
    Currency{"Sudanese Pound", "SDG", "SUDAN (THE)", "938", 2, ""}:                                                                       "SDG",
    Currency{"Surinam Dollar", "SRD", "SURINAME", "968", 2, "\u0024"}:                                                                    "SRD",
    Currency{"Lilangeni", "SZL", "ESWATINI", "748", 2, ""}:                                                                               "SZL",
    Currency{"Swedish Krona", "SEK", "SWEDEN", "752", 2, "\u006b\u0072"}:                                                                 "SEK",
    Currency{"WIR Euro", "CHE", "SWITZERLAND", "947", 2, ""}:                                                                             "CHE",
    Currency{"WIR Franc", "CHW", "SWITZERLAND", "948", 2, ""}:                                                                            "CHW",
    Currency{"Syrian Pound", "SYP", "SYRIAN ARAB REPUBLIC", "760", 2, "\u00a3"}:                                                          "SYP",
    Currency{"New Taiwan Dollar", "TWD", "TAIWAN (PROVINCE OF CHINA)", "901", 2, "\u004e\u0054\u0024"}:                                   "TWD",
    Currency{"Somoni", "TJS", "TAJIKISTAN", "972", 2, ""}:                                                                                "TJS",
    Currency{"Tanzanian Shilling", "TZS", "TANZANIA, UNITED REPUBLIC OF", "834", 2, ""}:                                                  "TZS",
    Currency{"Baht", "THB", "THAILAND", "764", 2, "\u0e3f"}:                                                                              "THB",
    Currency{"Pa’anga", "TOP", "TONGA", "776", 2, ""}:                                                                                    "TOP",
    Currency{"Trinidad and Tobago Dollar", "TTD", "TRINIDAD AND TOBAGO", "780", 2, "\u0054\u0054\u0024"}:                                 "TTD",
    Currency{"Tunisian Dinar", "TND", "TUNISIA", "788", 3, ""}:                                                                           "TND",
    Currency{"Turkish Lira", "TRY", "TURKEY", "949", 2, ""}:                                                                              "TRY",
    Currency{"Turkmenistan New Manat", "TMT", "TURKMENISTAN", "934", 2, ""}:                                                              "TMT",
    Currency{"Uganda Shilling", "UGX", "UGANDA", "800", 0, ""}:                                                                           "UGX",
    Currency{"Hryvnia", "UAH", "UKRAINE", "980", 2, "\u20b4"}:                                                                            "UAH",
    Currency{"UAE Dirham", "AED", "UNITED ARAB EMIRATES (THE)", "784", 2, ""}:                                                            "AED",
    Currency{"US Dollar (Next day)", "USN", "UNITED STATES OF AMERICA (THE)", "997", 2, ""}:                                              "USN",
    Currency{"Peso Uruguayo", "UYU", "URUGUAY", "858", 2, "\u0024\u0055"}:                                                                "UYU",
    Currency{"Uruguay Peso en Unidades Indexadas (UI)", "UYI", "URUGUAY", "940", 0, ""}:                                                  "UYI",
    Currency{"Unidad Previsional", "UYW", "URUGUAY", "927", 4, ""}:                                                                       "UYW",
    Currency{"Uzbekistan Sum", "UZS", "UZBEKISTAN", "860", 2, "\u043b\u0432"}:                                                            "UZS",
    Currency{"Vatu", "VUV", "VANUATU", "548", 0, ""}:                                                                                     "VUV",
    Currency{"Bolívar Soberano", "VES", "VENEZUELA (BOLIVARIAN REPUBLIC OF)", "928", 2, ""}:                                              "VES",
    Currency{"Dong", "VND", "VIET NAM", "704", 0, "\u20ab"}:                                                                              "VND",
    Currency{"Yemeni Rial", "YER", "YEMEN", "886", 2, "\ufdfc"}:                                                                          "YER",
    Currency{"Zambian Kwacha", "ZMW", "ZAMBIA", "967", 2, ""}:                                                                            "ZMW",
    Currency{"Zimbabwe Dollar", "ZWL", "ZIMBABWE", "932", 2, ""}:                                                                         "ZWL",
    Currency{"Bond Markets Unit European Composite Unit (EURCO)", "XBA", "ZZ01_Bond Markets Unit European_EURCO", "955", 0, ""}:          "XBA",
    Currency{"Bond Markets Unit European Monetary Unit (E.M.U.-6)", "XBB", "ZZ02_Bond Markets Unit European_EMU-6", "956", 0, ""}:        "XBB",
    Currency{"Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", "XBC", "ZZ03_Bond Markets Unit European_EUA-9", "957", 0, ""}:    "XBC",
    Currency{"Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", "XBD", "ZZ04_Bond Markets Unit European_EUA-17", "958", 0, ""}: "XBD",
    Currency{"Codes specifically reserved for testing purposes", "XTS", "ZZ06_Testing_Code", "963", 0, ""}:                               "XTS",
    Currency{"The codes assigned for transactions where no currency is involved", "XXX", "ZZ07_No_Currency", "999", 0, ""}:               "XXX",
    Currency{"Gold", "XAU", "ZZ08_Gold", "959", 0, ""}:                                                                                   "XAU",
    Currency{"Palladium", "XPD", "ZZ09_Palladium", "964", 0, ""}:                                                                         "XPD",
    Currency{"Platinum", "XPT", "ZZ10_Platinum", "962", 0, ""}:                                                                           "XPT",
    Currency{"Silver", "XAG", "ZZ11_Silver", "961", 0, ""}:                                                                               "XAG",
}
```
CurrencyName is a map Currency -> string (currency code)

``` go
var NameCurrency = map[string]Currency{

    "AFN": {"Afghani", "AFN", "AFGHANISTAN", "971", 2, "\u060b"},

    "EUR": {"Euro", "EUR", "ÅLAND ISLANDS", "978", 2, "\u20ac"},

    "ALL": {"Lek", "ALL", "ALBANIA", "008", 2, "\u004c\u0065\u006b"},

    "DZD": {"Algerian Dinar", "DZD", "ALGERIA", "012", 2, ""},

    "USD": {"US Dollar", "USD", "AMERICAN SAMOA", "840", 2, "\u0024"},

    "AOA": {"Kwanza", "AOA", "ANGOLA", "973", 2, ""},

    "XCD": {"East Caribbean Dollar", "XCD", "ANGUILLA", "951", 2, "\u0024"},

    "ARS": {"Argentine Peso", "ARS", "ARGENTINA", "032", 2, "\u0024"},

    "AMD": {"Armenian Dram", "AMD", "ARMENIA", "051", 2, ""},

    "AWG": {"Aruban Florin", "AWG", "ARUBA", "533", 2, "\u0192"},

    "AUD": {"Australian Dollar", "AUD", "AUSTRALIA", "036", 2, "\u0024"},

    "AZN": {"Azerbaijan Manat", "AZN", "AZERBAIJAN", "944", 2, "\u20bc"},

    "BSD": {"Bahamian Dollar", "BSD", "BAHAMAS (THE)", "044", 2, "\u0024"},

    "BHD": {"Bahraini Dinar", "BHD", "BAHRAIN", "048", 3, ""},

    "BDT": {"Taka", "BDT", "BANGLADESH", "050", 2, ""},

    "BBD": {"Barbados Dollar", "BBD", "BARBADOS", "052", 2, "\u0024"},

    "BYN": {"Belarusian Ruble", "BYN", "BELARUS", "933", 2, "\u0042\u0072"},

    "BZD": {"Belize Dollar", "BZD", "BELIZE", "084", 2, "\u0042\u005a\u0024"},

    "XOF": {"CFA Franc BCEAO", "XOF", "BENIN", "952", 0, ""},

    "BMD": {"Bermudian Dollar", "BMD", "BERMUDA", "060", 2, "\u0024"},

    "INR": {"Indian Rupee", "INR", "BHUTAN", "356", 2, ""},

    "BTN": {"Ngultrum", "BTN", "BHUTAN", "064", 2, ""},

    "BOB": {"Boliviano", "BOB", "BOLIVIA (PLURINATIONAL STATE OF)", "068", 2, "\u0024\u0062"},

    "BOV": {"Mvdol", "BOV", "BOLIVIA (PLURINATIONAL STATE OF)", "984", 2, ""},

    "BAM": {"Convertible Mark", "BAM", "BOSNIA AND HERZEGOVINA", "977", 2, "\u004b\u004d"},

    "BWP": {"Pula", "BWP", "BOTSWANA", "072", 2, "\u0050"},

    "NOK": {"Norwegian Krone", "NOK", "BOUVET ISLAND", "578", 2, "\u006b\u0072"},

    "BRL": {"Brazilian Real", "BRL", "BRAZIL", "986", 2, "\u0052\u0024"},

    "BND": {"Brunei Dollar", "BND", "BRUNEI DARUSSALAM", "096", 2, "\u0024"},

    "BGN": {"Bulgarian Lev", "BGN", "BULGARIA", "975", 2, "\u043b\u0432"},

    "BIF": {"Burundi Franc", "BIF", "BURUNDI", "108", 0, ""},

    "CVE": {"Cabo Verde Escudo", "CVE", "CABO VERDE", "132", 2, ""},

    "KHR": {"Riel", "KHR", "CAMBODIA", "116", 2, "\u17db"},

    "XAF": {"CFA Franc BEAC", "XAF", "CAMEROON", "950", 0, ""},

    "CAD": {"Canadian Dollar", "CAD", "CANADA", "124", 2, "\u0024"},

    "KYD": {"Cayman Islands Dollar", "KYD", "CAYMAN ISLANDS (THE)", "136", 2, "\u0024"},

    "CLP": {"Chilean Peso", "CLP", "CHILE", "152", 0, "\u0024"},

    "CLF": {"Unidad de Fomento", "CLF", "CHILE", "990", 4, ""},

    "CNY": {"Yuan Renminbi", "CNY", "CHINA", "156", 2, "\u00a5"},

    "COP": {"Colombian Peso", "COP", "COLOMBIA", "170", 2, "\u0024"},

    "COU": {"Unidad de Valor Real", "COU", "COLOMBIA", "970", 2, ""},

    "KMF": {"Comorian Franc ", "KMF", "COMOROS (THE)", "174", 0, ""},

    "CDF": {"Congolese Franc", "CDF", "CONGO (THE DEMOCRATIC REPUBLIC OF THE)", "976", 2, ""},

    "NZD": {"New Zealand Dollar", "NZD", "COOK ISLANDS (THE)", "554", 2, "\u0024"},

    "CRC": {"Costa Rican Colon", "CRC", "COSTA RICA", "188", 2, "\u20a1"},

    "HRK": {"Kuna", "HRK", "CROATIA", "191", 2, "\u006b\u006e"},

    "CUP": {"Cuban Peso", "CUP", "CUBA", "192", 2, "\u20b1"},

    "CUC": {"Peso Convertible", "CUC", "CUBA", "931", 2, ""},

    "ANG": {"Netherlands Antillean Guilder", "ANG", "CURAÇAO", "532", 2, "\u0192"},

    "CZK": {"Czech Koruna", "CZK", "CZECHIA", "203", 2, "\u004b\u010d"},

    "DKK": {"Danish Krone", "DKK", "DENMARK", "208", 2, "\u006b\u0072"},

    "DJF": {"Djibouti Franc", "DJF", "DJIBOUTI", "262", 0, ""},

    "DOP": {"Dominican Peso", "DOP", "DOMINICAN REPUBLIC (THE)", "214", 2, "\u0052\u0044\u0024"},

    "EGP": {"Egyptian Pound", "EGP", "EGYPT", "818", 2, "\u00a3"},

    "SVC": {"El Salvador Colon", "SVC", "EL SALVADOR", "222", 2, "\u0024"},

    "ERN": {"Nakfa", "ERN", "ERITREA", "232", 2, ""},

    "ETB": {"Ethiopian Birr", "ETB", "ETHIOPIA", "230", 2, ""},

    "FKP": {"Falkland Islands Pound", "FKP", "FALKLAND ISLANDS (THE) [MALVINAS]", "238", 2, "\u00a3"},

    "FJD": {"Fiji Dollar", "FJD", "FIJI", "242", 2, "\u0024"},

    "XPF": {"CFP Franc", "XPF", "FRENCH POLYNESIA", "953", 0, ""},

    "GMD": {"Dalasi", "GMD", "GAMBIA (THE)", "270", 2, ""},

    "GEL": {"Lari", "GEL", "GEORGIA", "981", 2, ""},

    "GHS": {"Ghana Cedi", "GHS", "GHANA", "936", 2, "\u00a2"},

    "GIP": {"Gibraltar Pound", "GIP", "GIBRALTAR", "292", 2, "\u00a3"},

    "GTQ": {"Quetzal", "GTQ", "GUATEMALA", "320", 2, "\u0051"},

    "GBP": {"Pound Sterling", "GBP", "GUERNSEY", "826", 2, "\u00a3"},

    "GNF": {"Guinean Franc", "GNF", "GUINEA", "324", 0, ""},

    "GYD": {"Guyana Dollar", "GYD", "GUYANA", "328", 2, "\u0024"},

    "HTG": {"Gourde", "HTG", "HAITI", "332", 2, ""},

    "HNL": {"Lempira", "HNL", "HONDURAS", "340", 2, "\u004c"},

    "HKD": {"Hong Kong Dollar", "HKD", "HONG KONG", "344", 2, "\u0024"},

    "HUF": {"Forint", "HUF", "HUNGARY", "348", 2, "\u0046\u0074"},

    "ISK": {"Iceland Krona", "ISK", "ICELAND", "352", 0, "\u006b\u0072"},

    "IDR": {"Rupiah", "IDR", "INDONESIA", "360", 2, "\u0052\u0070"},

    "XDR": {"SDR (Special Drawing Right)", "XDR", "INTERNATIONAL MONETARY FUND (IMF) ", "960", 0, ""},

    "IRR": {"Iranian Rial", "IRR", "IRAN (ISLAMIC REPUBLIC OF)", "364", 2, "\ufdfc"},

    "IQD": {"Iraqi Dinar", "IQD", "IRAQ", "368", 3, ""},

    "ILS": {"New Israeli Sheqel", "ILS", "ISRAEL", "376", 2, "\u20aa"},

    "JMD": {"Jamaican Dollar", "JMD", "JAMAICA", "388", 2, "\u004a\u0024"},

    "JPY": {"Yen", "JPY", "JAPAN", "392", 0, "\u00a5"},

    "JOD": {"Jordanian Dinar", "JOD", "JORDAN", "400", 3, ""},

    "KZT": {"Tenge", "KZT", "KAZAKHSTAN", "398", 2, "\u043b\u0432"},

    "KES": {"Kenyan Shilling", "KES", "KENYA", "404", 2, ""},

    "KPW": {"North Korean Won", "KPW", "KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)", "408", 2, "\u20a9"},

    "KRW": {"Won", "KRW", "KOREA (THE REPUBLIC OF)", "410", 0, "\u20a9"},

    "KWD": {"Kuwaiti Dinar", "KWD", "KUWAIT", "414", 3, ""},

    "KGS": {"Som", "KGS", "KYRGYZSTAN", "417", 2, "\u043b\u0432"},

    "LAK": {"Lao Kip", "LAK", "LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)", "418", 2, "\u20ad"},

    "LBP": {"Lebanese Pound", "LBP", "LEBANON", "422", 2, "\u00a3"},

    "LSL": {"Loti", "LSL", "LESOTHO", "426", 2, ""},

    "ZAR": {"Rand", "ZAR", "LESOTHO", "710", 2, "\u0052"},

    "LRD": {"Liberian Dollar", "LRD", "LIBERIA", "430", 2, "\u0024"},

    "LYD": {"Libyan Dinar", "LYD", "LIBYA", "434", 3, ""},

    "CHF": {"Swiss Franc", "CHF", "LIECHTENSTEIN", "756", 2, "\u0043\u0048\u0046"},

    "MOP": {"Pataca", "MOP", "MACAO", "446", 2, ""},

    "MKD": {"Denar", "MKD", "MACEDONIA (THE FORMER YUGOSLAV REPUBLIC OF)", "807", 2, "\u0434\u0435\u043d"},

    "MGA": {"Malagasy Ariary", "MGA", "MADAGASCAR", "969", 2, ""},

    "MWK": {"Malawi Kwacha", "MWK", "MALAWI", "454", 2, ""},

    "MYR": {"Malaysian Ringgit", "MYR", "MALAYSIA", "458", 2, "\u0052\u004d"},

    "MVR": {"Rufiyaa", "MVR", "MALDIVES", "462", 2, ""},

    "MRU": {"Ouguiya", "MRU", "MAURITANIA", "929", 2, ""},

    "MUR": {"Mauritius Rupee", "MUR", "MAURITIUS", "480", 2, "\u20a8"},

    "XUA": {"ADB Unit of Account", "XUA", "MEMBER COUNTRIES OF THE AFRICAN DEVELOPMENT BANK GROUP", "965", 0, ""},

    "MXN": {"Mexican Peso", "MXN", "MEXICO", "484", 2, "\u0024"},

    "MXV": {"Mexican Unidad de Inversion (UDI)", "MXV", "MEXICO", "979", 2, ""},

    "MDL": {"Moldovan Leu", "MDL", "MOLDOVA (THE REPUBLIC OF)", "498", 2, ""},

    "MNT": {"Tugrik", "MNT", "MONGOLIA", "496", 2, "\u20ae"},

    "MAD": {"Moroccan Dirham", "MAD", "MOROCCO", "504", 2, ""},

    "MZN": {"Mozambique Metical", "MZN", "MOZAMBIQUE", "943", 2, "\u004d\u0054"},

    "MMK": {"Kyat", "MMK", "MYANMAR", "104", 2, ""},

    "NAD": {"Namibia Dollar", "NAD", "NAMIBIA", "516", 2, "\u0024"},

    "NPR": {"Nepalese Rupee", "NPR", "NEPAL", "524", 2, "\u20a8"},

    "NIO": {"Cordoba Oro", "NIO", "NICARAGUA", "558", 2, "\u0043\u0024"},

    "NGN": {"Naira", "NGN", "NIGERIA", "566", 2, "\u20a6"},

    "OMR": {"Rial Omani", "OMR", "OMAN", "512", 3, "\ufdfc"},

    "PKR": {"Pakistan Rupee", "PKR", "PAKISTAN", "586", 2, "\u20a8"},

    "PAB": {"Balboa", "PAB", "PANAMA", "590", 2, "\u0042\u002f\u002e"},

    "PGK": {"Kina", "PGK", "PAPUA NEW GUINEA", "598", 2, ""},

    "PYG": {"Guarani", "PYG", "PARAGUAY", "600", 0, "\u0047\u0073"},

    "PEN": {"Sol", "PEN", "PERU", "604", 2, "\u0053\u002f\u002e"},

    "PHP": {"Philippine Peso", "PHP", "PHILIPPINES (THE)", "608", 2, "\u20b1"},

    "PLN": {"Zloty", "PLN", "POLAND", "985", 2, "\u007a\u0142"},

    "QAR": {"Qatari Rial", "QAR", "QATAR", "634", 2, "\ufdfc"},

    "RON": {"Romanian Leu", "RON", "ROMANIA", "946", 2, "\u006c\u0065\u0069"},

    "RUB": {"Russian Ruble", "RUB", "RUSSIAN FEDERATION (THE)", "643", 2, "\u20bd"},

    "RWF": {"Rwanda Franc", "RWF", "RWANDA", "646", 0, ""},

    "SHP": {"Saint Helena Pound", "SHP", "SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA", "654", 2, "\u00a3"},

    "WST": {"Tala", "WST", "SAMOA", "882", 2, ""},

    "STN": {"Dobra", "STN", "SAO TOME AND PRINCIPE", "930", 2, ""},

    "SAR": {"Saudi Riyal", "SAR", "SAUDI ARABIA", "682", 2, "\ufdfc"},

    "RSD": {"Serbian Dinar", "RSD", "SERBIA", "941", 2, "\u0414\u0438\u043d\u002e"},

    "SCR": {"Seychelles Rupee", "SCR", "SEYCHELLES", "690", 2, "\u20a8"},

    "SLL": {"Leone", "SLL", "SIERRA LEONE", "694", 2, ""},

    "SGD": {"Singapore Dollar", "SGD", "SINGAPORE", "702", 2, "\u0024"},

    "XSU": {"Sucre", "XSU", "SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS \"SUCRE\"", "994", 0, ""},

    "SBD": {"Solomon Islands Dollar", "SBD", "SOLOMON ISLANDS", "090", 2, "\u0024"},

    "SOS": {"Somali Shilling", "SOS", "SOMALIA", "706", 2, "\u0053"},

    "SSP": {"South Sudanese Pound", "SSP", "SOUTH SUDAN", "728", 2, ""},

    "LKR": {"Sri Lanka Rupee", "LKR", "SRI LANKA", "144", 2, "\u20a8"},

    "SDG": {"Sudanese Pound", "SDG", "SUDAN (THE)", "938", 2, ""},

    "SRD": {"Surinam Dollar", "SRD", "SURINAME", "968", 2, "\u0024"},

    "SZL": {"Lilangeni", "SZL", "ESWATINI", "748", 2, ""},

    "SEK": {"Swedish Krona", "SEK", "SWEDEN", "752", 2, "\u006b\u0072"},

    "CHE": {"WIR Euro", "CHE", "SWITZERLAND", "947", 2, ""},

    "CHW": {"WIR Franc", "CHW", "SWITZERLAND", "948", 2, ""},

    "SYP": {"Syrian Pound", "SYP", "SYRIAN ARAB REPUBLIC", "760", 2, "\u00a3"},

    "TWD": {"New Taiwan Dollar", "TWD", "TAIWAN (PROVINCE OF CHINA)", "901", 2, "\u004e\u0054\u0024"},

    "TJS": {"Somoni", "TJS", "TAJIKISTAN", "972", 2, ""},

    "TZS": {"Tanzanian Shilling", "TZS", "TANZANIA, UNITED REPUBLIC OF", "834", 2, ""},

    "THB": {"Baht", "THB", "THAILAND", "764", 2, "\u0e3f"},

    "TOP": {"Pa’anga", "TOP", "TONGA", "776", 2, ""},

    "TTD": {"Trinidad and Tobago Dollar", "TTD", "TRINIDAD AND TOBAGO", "780", 2, "\u0054\u0054\u0024"},

    "TND": {"Tunisian Dinar", "TND", "TUNISIA", "788", 3, ""},

    "TRY": {"Turkish Lira", "TRY", "TURKEY", "949", 2, ""},

    "TMT": {"Turkmenistan New Manat", "TMT", "TURKMENISTAN", "934", 2, ""},

    "UGX": {"Uganda Shilling", "UGX", "UGANDA", "800", 0, ""},

    "UAH": {"Hryvnia", "UAH", "UKRAINE", "980", 2, "\u20b4"},

    "AED": {"UAE Dirham", "AED", "UNITED ARAB EMIRATES (THE)", "784", 2, ""},

    "USN": {"US Dollar (Next day)", "USN", "UNITED STATES OF AMERICA (THE)", "997", 2, ""},

    "UYU": {"Peso Uruguayo", "UYU", "URUGUAY", "858", 2, "\u0024\u0055"},

    "UYI": {"Uruguay Peso en Unidades Indexadas (UI)", "UYI", "URUGUAY", "940", 0, ""},

    "UYW": {"Unidad Previsional", "UYW", "URUGUAY", "927", 4, ""},

    "UZS": {"Uzbekistan Sum", "UZS", "UZBEKISTAN", "860", 2, "\u043b\u0432"},

    "VUV": {"Vatu", "VUV", "VANUATU", "548", 0, ""},

    "VES": {"Bolívar Soberano", "VES", "VENEZUELA (BOLIVARIAN REPUBLIC OF)", "928", 2, ""},

    "VND": {"Dong", "VND", "VIET NAM", "704", 0, "\u20ab"},

    "YER": {"Yemeni Rial", "YER", "YEMEN", "886", 2, "\ufdfc"},

    "ZMW": {"Zambian Kwacha", "ZMW", "ZAMBIA", "967", 2, ""},

    "ZWL": {"Zimbabwe Dollar", "ZWL", "ZIMBABWE", "932", 2, ""},

    "XBA": {"Bond Markets Unit European Composite Unit (EURCO)", "XBA", "ZZ01_Bond Markets Unit European_EURCO", "955", 0, ""},

    "XBB": {"Bond Markets Unit European Monetary Unit (E.M.U.-6)", "XBB", "ZZ02_Bond Markets Unit European_EMU-6", "956", 0, ""},

    "XBC": {"Bond Markets Unit European Unit of Account 9 (E.U.A.-9)", "XBC", "ZZ03_Bond Markets Unit European_EUA-9", "957", 0, ""},

    "XBD": {"Bond Markets Unit European Unit of Account 17 (E.U.A.-17)", "XBD", "ZZ04_Bond Markets Unit European_EUA-17", "958", 0, ""},

    "XTS": {"Codes specifically reserved for testing purposes", "XTS", "ZZ06_Testing_Code", "963", 0, ""},

    "XXX": {"The codes assigned for transactions where no currency is involved", "XXX", "ZZ07_No_Currency", "999", 0, ""},

    "XAU": {"Gold", "XAU", "ZZ08_Gold", "959", 0, ""},

    "XPD": {"Palladium", "XPD", "ZZ09_Palladium", "964", 0, ""},

    "XPT": {"Platinum", "XPT", "ZZ10_Platinum", "962", 0, ""},

    "XAG": {"Silver", "XAG", "ZZ11_Silver", "961", 0, ""},
}
```
NameCurrency is a map string(currency code) -> Currency




## <a name="Currency">type</a> [Currency](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/currency/currency.go?s=195:420#L11)
``` go
type Currency struct {
    Name       string `json:"name"`
    Code       string `json:"code"`
    Country    string `json:"country"`
    NumberCode string `json:"-"`
    Units      uint8  `json:"units"`
    Symbol     string `json:"symbol"`
}

```
Currency is an enum over list of currencies presented at ISO 4217







### <a name="Parse">func</a> [Parse](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/currency/currency.go?s=47667:47705#L566)
``` go
func Parse(n string) (Currency, error)
```
Parse string to Currency





### <a name="Currency.MarshalJSON">func</a> (Currency) [MarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/currency/currency.go?s=48184:48231#L589)
``` go
func (c Currency) MarshalJSON() ([]byte, error)
```
MarshalJSON implements json.Marshaler interface




### <a name="Currency.UnmarshalJSON">func</a> (\*Currency) [UnmarshalJSON](https://git.elewise.com/elma365/common/-/tree/develop/pkg/types/complextypes/money/currency/currency.go?s=47894:47945#L574)
``` go
func (c *Currency) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Unmarshaler interface







- - -
Generated by [godoc2md](https://github.com/Exa-Networks/godoc2md)
