#!/usr/bin/env python3
# -*- coding: utf8 -*-
'''
currency_list_update is a simple script to retrieve currency list
from currency-iso.org as xml, parse it and generate currency.go
'''
import requests
from lxml import html
from urllib.request import urlopen
from xml.etree import ElementTree


class Currency:
    def __init__(self, country, name, code, numberCode, units, symbol):
        self.country = country
        self.name = name
        self.code = code
        self.numberCode = numberCode
        self.units = units
        self.symbol = symbol


URL = 'https://www.currency-iso.org/dam/downloads/lists/list_one.xml'
SymbolsURL = 'https://www.xe.com/symbols.php'

if __name__ == '__main__':
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:45.0) Gecko/20100101 Firefox/45.0'
    }
    r = requests.get(SymbolsURL, headers=headers)
    tree = html.fromstring(r.text)
    currency_list_lxml = tree.xpath('//table[@class = "currencySymblTable"]')[0]
    rows = currency_list_lxml.xpath('//tr[@class != "currencySymblTableSubTitle"]')

    symbolsMap = {}

    for row in rows:
        currencyCode = row.xpath('./td')[1].text
        unicodeHex = row.xpath('./td')[6].text
        if unicodeHex is None:
            continue
        hexiarr = unicodeHex.split(',')
        result = b''
        if len(hexiarr) > 0:
            for st in hexiarr:
                st = st.strip()
                while len(st) < 4:
                    st = '0' + st
                st = b'\u' + st.encode('utf-8')
                result = result + st

        symbolsMap[currencyCode] = result

    src = urlopen(URL).read()
    root = ElementTree.fromstring(src)
    currencyDict = {
        'ATS': Currency('Austria', 'Austrian schilling', 'ATS', 978, 2, 'ATS'),
        'BEF': Currency('Belgium', 'Belgian franc', 'BEF', 978, 2, 'BEF'),
        'DEM': Currency('Germany', 'Deutsche Mark', 'DEM', 978, 2, 'DEM'),
        'EEK': Currency('Estonia', 'Estonian kroon', 'EEK', 978, 2, 'EEK'),
        'ESP': Currency('Spain', 'Spanish peseta', 'ESP', 978, 2, 'ESP'),
        'FIM': Currency('Finland', 'Finnish mark', 'FIM', 978, 2, 'FIM'),
        'FRF': Currency('France', 'French franc', 'FRF', 978, 2, 'FRF'),
        'GRD': Currency('Greece', 'Greek drachma', 'GRD', 978, 2, 'GRD'),
        'IEP': Currency('Ireland', 'Irish pound', 'IEP', 978, 2, 'IEP'),
        'ITL': Currency('Italy', 'Italian lira', 'ITL', 978, 2, 'ITL'),
        'LUF': Currency('Luxembourg', 'Luxembourg franc', 'LUF', 978, 2, 'LUF'),
        'NLG': Currency('Netherlands', 'Netherlands guilder', 'NLG', 978, 2, 'NLG'),
        'PTE': Currency('Portugal', 'Portuguese escudo', 'PTE', 978, 2, 'PTE'),
        'GHC': Currency('Ghana', 'Ghanaian cedi', 'GHC', 288, 2, 'GHC'),
        'GWP': Currency('Guinea-Bissau', 'Guinea-Bissau peso', 'GWP', 598, 2, 'GWP'),
        'LVL': Currency('Latvia', 'Latvian Lats', 'LVL', 428, 2, 'LVL'),
        'LTL': Currency('Lithuania', 'Lithuanian litas', 'LTL', 440, 2, 'LTL'),
        'MTL': Currency('Malta', 'Maltese lira', 'MTL', 470, 2, 'MTL'),
        'MRO': Currency('Mauritania', 'Mauritanian ouguiya', 'MRO', 478, 2, 'MRO'),
        'ROL': Currency('Romania', 'Romanian leu', 'ROL', 642, 2, 'ROL'),
        'MZM': Currency('Mozambique', 'Mozambican metical', 'MZM', 508, 2, 'MZM'),
        'SAC': Currency('South Africa', 'South African rand', 'SAC', 710, 2, 'SAC'),
        'SIT': Currency('Slovenia', 'Slovenian tolar', 'SIT', 0, 2, 'SIT'),
        'SDD': Currency('Sudan', 'Sudanese dinar', 'SDD', 705, 2, 'SDD'),
        'TMM': Currency('Turkmenistan', 'Turkmenistani manat', 'TMM', 795, 2, 'TMM'),
        'ZMK': Currency('Zambia', 'Zambian kwacha', 'ZMK', 894, 2, 'ZMK'),
        'ZWD': Currency('Zimbabwe', 'Zimbabwean dollar', 'ZWD', 716, 2, 'ZWD'),
    }

    m = {}
    for entry in root.iter('CcyNtry'):
        country = entry.find('CtryNm')
        name = entry.find('CcyNm')
        code = entry.find('Ccy')
        numberCode = entry.find('CcyNbr')
        units = entry.find('CcyMnrUnts')
        if code is None:
            continue

        if code.text in m:
            continue

        m[code.text] = True
        symbol = symbolsMap.get(code.text)
        if symbol is None:
            symbol = b''
        newCurrency = Currency(country.text.replace('"', '\\"'), name.text, code.text, numberCode.text,
                               0 if units.text == 'N.A.' else units.text, symbol.decode('utf-8'))

        currencyDict[code.text] = newCurrency

    print('// Code generated by currency_list_update.')
    print('// DO NOT EDIT!')
    print()
    print('package currency')
    print()
    print('import "fmt"')
    print('import "strconv"')
    print('import "strings"')
    print()
    print('// Currency is an enum over list of currencies presented at ISO 4217')
    print('type Currency struct {')
    print('\tName string `json:"name"`')
    print('\tCode string `json:"code"`')
    print('\tCountry string `json:"country"`')
    print('\tNumberCode string `json:"-"`')
    print('\tUnits uint8 `json:"units"`')
    print('\tSymbol string `json:"symbol"`')
    print('}')
    print()

    print('// CurrencyName is a map Currency -> string (currency code)')
    print('var CurrencyName = map[Currency]string{')
    for code in currencyDict:
        cur = currencyDict[code]
        currencyObject = '"{name}", "{code}", "{country}", "{numberCode}", {units}, "{symbol}"'.format(name=cur.name,
                                                                                                       code=cur.code,
                                                                                                       country=cur.country,
                                                                                                       numberCode=cur.numberCode,
                                                                                                       units=cur.units,
                                                                                                       symbol=cur.symbol)
        print('\tCurrency{' + currencyObject + '}:\t"' + cur.code + '",')
    print('}')
    print()

    print('// NameCurrency is a map string(currency code) -> Currency')
    print('var NameCurrency = map[string]Currency{')
    for code in currencyDict:
        cur = currencyDict[code]
        comment = '{ccy} is the {name} used in {country}'.format(ccy=cur.code, name=cur.name, country=cur.country)
        currencyObject = '"{name}", "{code}", "{country}", "{numberCode}", {units}, "{symbol}"'.format(name=cur.name,
                                                                                                       code=cur.code,
                                                                                                       country=cur.country,
                                                                                                       numberCode=cur.numberCode,
                                                                                                       units=cur.units,
                                                                                                       symbol=cur.symbol)
        print('\t// ' + comment)
        print('\t"' + cur.code + '":\t{' + currencyObject + '},')
    print('}')
    print()
    print('// Parse string to Currency')
    print('func Parse(n string) (Currency, error) {')
    print('\tif c, ok := NameCurrency[strings.ToUpper(n)]; ok {')
    print('\t\treturn c, nil')
    print('\t}')
    print('\treturn Currency{}, fmt.Errorf("unknown currency %q", n)')
    print('}')
    print()
    print('// UnmarshalJSON implements json.Unmarshaler interface')
    print('func (c *Currency) UnmarshalJSON(data []byte) error {')
    print('\tsdata := string(data)')
    print('\tif len(sdata) > 1 && sdata[0] == \'"\' {')
    print('\t\tsdata = sdata[1 : len(sdata)-1]')
    print('\t}')
    print('\tcur, err := Parse(sdata)')
    print('\tif err != nil {')
    print('\t\treturn err')
    print('\t}')
    print('\t*c = cur')
    print('')
    print('\treturn nil')
    print('}')
    print()
    print('// MarshalJSON implements json.Marshaler interface')
    print('func (c Currency) MarshalJSON() ([]byte, error) {')
    print('')
    print('\treturn []byte(strconv.Quote(c.Code)), nil')
    print('}')
