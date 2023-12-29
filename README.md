# event-tracking

Event tracking microservice for the social networking app focused on sharing vibes

```bash
docker build -t event-tracking .
docker tag event-tracking:latest vibesocial/event-tracking:latest
docker push vibesocial/event-tracking:latest
```

TODO:

- Zvezi metrike z prometheusom in azure dashboardom
- Spisi dokumentacijo
- Dodaj Ingress controllre (in vse potrebno)
- V eno mikrostoritev vključite izolacijo in toleranco napak. Pripravite demonstracijo mehanizmov na primeru. Ocenjuje se tudi razumevanje primera, ki ste ga vključili v vašo rešitev. (12T)
- Vključite centralizirano beleženja dnevnikov. Za vsakega člana skupine pripravite en primer zanimive poizvedbe po dnevnikih. Nadalje še demonstrirajte sledenje zahtevkov pri obdelavi na različnih mikrostoritvah. (12T)
- premakni okoljske spremenljivke v kubernetes konfiguracijo
- naredi zvezno integracijo
- poglej za konfiguracijski streznik
- poglej za auth service (nest js in supabase + auth0) (async rest)
