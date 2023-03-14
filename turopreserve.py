import sys
from urllib.parse import urlparse, parse_qs, quote, urlencode

def dedu(url):
    parsed_url = urlparse(url)
    query_dict = parse_qs(parsed_url.query, keep_blank_values=True)
    deduplicated_query_dict = {k: v[0] for k, v in query_dict.items()}

    new_url = parsed_url._replace(query=urlencode(deduplicated_query_dict, doseq=True)).geturl()

    print(new_url)


filename = sys.argv[1]

with open(filename) as f:
    urls = f.read().splitlines()

url_dict = {}

for url in urls:
    parsed_url = urlparse(url)
    if parsed_url.query:
        url_key = f"{parsed_url.scheme}://{parsed_url.netloc}{parsed_url.path}?"
        query_dict = parse_qs(parsed_url.query)
        url_value = []
        seen_params = set()
        for k, v_list in query_dict.items():
            for v in v_list:
                if k not in seen_params:
                    url_value.append(f"{k}={quote(v.encode())}")
                    seen_params.add(k)
        url_dict.setdefault(url_key, []).extend(url_value)

for url_key, url_value in url_dict.items():
    uralo=f"{url_key}{'&'.join(url_value)}"
    dedu(uralo)



