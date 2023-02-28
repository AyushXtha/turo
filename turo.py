import sys
from urllib.parse import urlparse, parse_qs


with open(sys.argv[1]) as fp:
    urls = filter(None, fp.read().split('\n'))

url_list = {}
for url in urls:
    o = urlparse(url)
    if o.query:
        urlkey = o.hostname + o.path + '?'
        urlvalue = set(parse_qs(o.query))
        url_list.setdefault(urlkey, set()).update(urlvalue)

for url, params in url_list.items():
    print('https://' + url + '=FUZZ&'.join(params) + '=FUZZ')
