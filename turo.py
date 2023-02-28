from urllib.parse import urlparse,parse_qs
import sys


file_name = sys.argv[1]
fp = open(file_name)    
contents = fp.read()
fp.close()

urls= list(filter(None,contents.split('\n')))

url_list={}
for url in urls:
    o = urlparse(url)
    if (o.query != ''):
        urlkey=o.hostname+o.path+'?'
        urlvalue=(list(parse_qs(o.query).keys()))  
        if urlkey in url_list:
            url_list[urlkey].extend(urlvalue)
        else:
            url_list[urlkey]=urlvalue


for url in url_list:
    print('https://'+url+'=FUZZ&'.join(set(url_list[url]))+'=FUZZ')


    
