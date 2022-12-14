# camdl-dns

Instruction:
1. Install go packages with

        go get github.com/ethereum/go-ethereum/cmd/devp2p

        go get github.com/ethereum/go-ethereum/cmd/ethkey

    For cmd/powershell, might need to run this instead

        go install <package>@latest 

<br>

2. Create key for signing DNS tree with

        bash genkey.sh 
    
    This will create `data/dnskey.json`.

<br>

3. Edit node info in `data/raw-nodes.json`. 
    - This will create `data/nodes.json`.
    - assign `v4` to `id` attribute. (other options haven't been considered)
    - `seq` should (maybe) increase when changing an existing node info.
    - `privKey` is node private key.

<br>

4. Create signed DNS tree and export TXT record with 
        
        bash export.sh <domain> <seq>
    
    - This will create `data/enrtree-info.json` via `devp2p dns sign` cmd.
    - This will also create `data/TXT.json` via `devp2p dns to-txt`.

<br>

5. Add records from `data/TXT.json` to DNS.

<br>

6. (Optinonal) to upload records to cloudflare with cloudflare API,
    
    6.1. Add env variables in a `.env` file as shown in `.env.example`

    - CLOUDFLARE_API_TOKEN
    - ZONE_ID

    <br>

    6.2 Publish record to cloudflare with 

        bash publish.sh

    cloudflare API doesn't support some TLDs such as .tk, .ml

    