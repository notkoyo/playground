// Rank Command

$(eval res=`$(urlfetch json https://api.henrikdev.xyz/valorant/v1/mmr/{region}/{name}/{tag}?api_key={API_KEY})`;try{json=JSON.parse(res);json.data.currenttierpatched && json.data.ranking_in_tier == null ? "Error fetching rank." : json.data.currenttierpatched + " - " + json.data.ranking_in_tier + " RR"; } catch(e){`${e}: ${res}`.substr(0, 400)})

// RR Today Command

$(eval res=`$(urlfetch json https://api.henrikdev.xyz/valorant/v1/mmr-history/{region}/{name}/{tag}?api_key={API_KEY})`;try{td=JSON.parse(res).data;rr=0;for(i=0;i<td.length; i++){if(td[i].date.includes(`$(time Europe/London "MMMM D")`)){rr=rr+td[i].mmr_change_to_last_game;}}rr<0 ? msg="{name} has lost " + rr*-1 + " RR today.":msg="{name} has gained " + rr + " RR today.";msg;} catch(e){`${e}: ${response}`.substr(0, 400)})