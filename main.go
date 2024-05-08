package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Emoji struct {
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

var emojis = []Emoji{
	{Name: "Grinning Face", Emoji: "😀"},
	{Name: "Grinning Face with Big Eyes", Emoji: "😃"},
	{Name: "Grinning Face with Smiling Eyes", Emoji: "😄"},
	{Name: "Beaming Face with Smiling Eyes", Emoji: "😁"},
	{Name: "Grinning Squinting Face", Emoji: "😆"},
	{Name: "Grinning Face with Sweat", Emoji: "😅"},
	{Name: "Rolling on the Floor Laughing", Emoji: "🤣"},
	{Name: "Face with Tears of Joy", Emoji: "😂"},
	{Name: "Slightly Smiling Face", Emoji: "🙂"},
	{Name: "Upside-Down Face", Emoji: "🙃"},
	{Name: "Winking Face", Emoji: "😉"},
	{Name: "Smiling Face with Smiling Eyes", Emoji: "😊"},
	{Name: "Smiling Face with Halo", Emoji: "😇"},
	{Name: "Smiling Face with Hearts", Emoji: "🥰"},
	{Name: "Smiling Face with Heart-Eyes", Emoji: "😍"},
	{Name: "Star-Struck", Emoji: "🤩"},
	{Name: "Face Blowing a Kiss", Emoji: "😘"},
	{Name: "Kissing Face", Emoji: "😗"},
	{Name: "Smiling Face", Emoji: "☺️"},
	{Name: "Kissing Face with Closed Eyes", Emoji: "😚"},
	{Name: "Kissing Face with Smiling Eyes", Emoji: "😙"},
	{Name: "Face Savoring Food", Emoji: "😋"},
	{Name: "Face with Tongue", Emoji: "😛"},
	{Name: "Winking Face with Tongue", Emoji: "😜"},
	{Name: "Zany Face", Emoji: "🤪"},
	{Name: "Squinting Face with Tongue", Emoji: "😝"},
	{Name: "Money-Mouth Face", Emoji: "🤑"},
	{Name: "Hugging Face", Emoji: "🤗"},
	{Name: "Face with Hand Over Mouth", Emoji: "🤭"},
	{Name: "Shushing Face", Emoji: "🤫"},
	{Name: "Thinking Face", Emoji: "🤔"},
	{Name: "Zipper-Mouth Face", Emoji: "🤐"},
	{Name: "Face with Raised Eyebrow", Emoji: "🤨"},
	{Name: "Neutral Face", Emoji: "😐"},
	{Name: "Expressionless Face", Emoji: "😑"},
	{Name: "Face Without Mouth", Emoji: "😶"},
	{Name: "Smirking Face", Emoji: "😏"},
	{Name: "Unamused Face", Emoji: "😒"},
	{Name: "Face with Rolling Eyes", Emoji: "🙄"},
	{Name: "Grimacing Face", Emoji: "😬"},
	{Name: "Lying Face", Emoji: "🤥"},
	{Name: "Relieved Face", Emoji: "😌"},
	{Name: "Pensive Face", Emoji: "😔"},
	{Name: "Sleepy Face", Emoji: "😪"},
	{Name: "Drooling Face", Emoji: "🤤"},
	{Name: "Sleeping Face", Emoji: "😴"},
	{Name: "Face with Medical Mask", Emoji: "😷"},
	{Name: "Face with Thermometer", Emoji: "🤒"},
	{Name: "Face with Head-Bandage", Emoji: "🤕"},
	{Name: "Nauseated Face", Emoji: "🤢"},
	{Name: "Face Vomiting", Emoji: "🤮"},
	{Name: "Sneezing Face", Emoji: "🤧"},
	{Name: "Hot Face", Emoji: "🥵"},
	{Name: "Cold Face", Emoji: "🥶"},
	{Name: "Woozy Face", Emoji: "🥴"},
	{Name: "Dizzy Face", Emoji: "😵"},
	{Name: "Exploding Head", Emoji: "🤯"},
	{Name: "Cowboy Hat Face", Emoji: "🤠"},
	{Name: "Partying Face", Emoji: "🥳"},
	{Name: "Disguised Face", Emoji: "🥸"},
	{Name: "Smiling Face with Sunglasses", Emoji: "😎"},
	{Name: "Nerd Face", Emoji: "🤓"},
	{Name: "Face with Monocle", Emoji: "🧐"},
	{Name: "Confused Face", Emoji: "😕"},
	{Name: "Worried Face", Emoji: "😟"},
	{Name: "Slightly Frowning Face", Emoji: "🙁"},
	{Name: "Frowning Face", Emoji: "☹️"},
	{Name: "Face with Open Mouth", Emoji: "😮"},
	{Name: "Hushed Face", Emoji: "😯"},
	{Name: "Astonished Face", Emoji: "😲"},
	{Name: "Flushed Face", Emoji: "😳"},
	{Name: "Pleading Face", Emoji: "🥺"},
	{Name: "Frowning Face with Open Mouth", Emoji: "😦"},
	{Name: "Anguished Face", Emoji: "😧"},
	{Name: "Fearful Face", Emoji: "😨"},
	{Name: "Anxious Face with Sweat", Emoji: "😰"},
	{Name: "Sad but Relieved Face", Emoji: "😥"},
	{Name: "Crying Face", Emoji: "😢"},
	{Name: "Loudly Crying Face", Emoji: "😭"},
	{Name: "Face Screaming in Fear", Emoji: "😱"},
	{Name: "Confounded Face", Emoji: "😖"},
	{Name: "Persevering Face", Emoji: "😣"},
	{Name: "Disappointed Face", Emoji: "😞"},
	{Name: "Downcast Face with Sweat", Emoji: "😓"},
	{Name: "Weary Face", Emoji: "😩"},
	{Name: "Tired Face", Emoji: "😫"},
	{Name: "Yawning Face", Emoji: "🥱"},
	{Name: "Face with Steam From Nose", Emoji: "😤"},
	{Name: "Angry Face", Emoji: "😠"},
	{Name: "Pouting Face", Emoji: "😡"},
	{Name: "Face with Symbols on Mouth", Emoji: "🤬"},
	{Name: "Smiling Face with Horns", Emoji: "😈"},
	{Name: "Angry Face with Horns", Emoji: "👿"},
	{Name: "Skull", Emoji: "💀"},
	{Name: "Skull and Crossbones", Emoji: "☠️"},
	{Name: "Pile of Poo", Emoji: "💩"},
	{Name: "Clown Face", Emoji: "🤡"},
	{Name: "Ogre", Emoji: "👹"},
	{Name: "Goblin", Emoji: "👺"},
	{Name: "Ghost", Emoji: "👻"},
	{Name: "Alien", Emoji: "👽"},
	{Name: "Alien Monster", Emoji: "👾"},
	{Name: "Robot", Emoji: "🤖"},
	{Name: "Grinning Cat", Emoji: "😺"},
	{Name: "Grinning Cat with Smiling Eyes", Emoji: "😸"},
	{Name: "Cat with Tears of Joy", Emoji: "😹"},
	{Name: "Smiling Cat with Heart-Eyes", Emoji: "😻"},
	{Name: "Cat with Wry Smile", Emoji: "😼"},
	{Name: "Kissing Cat", Emoji: "😽"},
	{Name: "Weary Cat", Emoji: "🙀"},
	{Name: "Crying Cat", Emoji: "😿"},
	{Name: "Pouting Cat", Emoji: "😾"},
	{Name: "See-No-Evil Monkey", Emoji: "🙈"},
	{Name: "Hear-No-Evil Monkey", Emoji: "🙉"},
	{Name: "Speak-No-Evil Monkey", Emoji: "🙊"},
	{Name: "Kiss Mark", Emoji: "💋"},
	{Name: "Love Letter", Emoji: "💌"},
	{Name: "Heart with Arrow", Emoji: "💘"},
	{Name: "Heart with Ribbon", Emoji: "💝"},
	{Name: "Sparkling Heart", Emoji: "💖"},
	{Name: "Growing Heart", Emoji: "💗"},
	{Name: "Beating Heart", Emoji: "💓"},
	{Name: "Revolving Hearts", Emoji: "💞"},
	{Name: "Two Hearts", Emoji: "💕"},
	{Name: "Heart Decoration", Emoji: "💟"},
	{Name: "Heart Exclamation", Emoji: "❣️"},
	{Name: "Broken Heart", Emoji: "💔"},
	{Name: "Red Heart", Emoji: "❤️"},
	{Name: "Orange Heart", Emoji: "🧡"},
	{Name: "Yellow Heart", Emoji: "💛"},
	{Name: "Green Heart", Emoji: "💚"},
	{Name: "Blue Heart", Emoji: "💙"},
	{Name: "Purple Heart", Emoji: "💜"},
	{Name: "Black Heart", Emoji: "🖤"},
	{Name: "Hundred Points", Emoji: "💯"},
	{Name: "Anger Symbol", Emoji: "💢"},
	{Name: "Collision", Emoji: "💥"},
	{Name: "Dizzy", Emoji: "💫"},
	{Name: "Speech Balloon", Emoji: "💬"},
	{Name: "Left Speech Bubble", Emoji: "🗨️"},
	{Name: "Right Anger Bubble", Emoji: "🗯️"},
	{Name: "Thought Balloon", Emoji: "💭"},
	{Name: "Zzz", Emoji: "💤"},
	{Name: "Splash", Emoji: "💦"},
	{Name: "Sweat Droplets", Emoji: "💧"},
	{Name: "Dash", Emoji: "💨"},
	{Name: "Hole", Emoji: "🕳️"},
	{Name: "Bomb", Emoji: "💣"},
	{Name: "Fire", Emoji: "🔥"},
	{Name: "Droplet", Emoji: "💧"},
	{Name: "Collision", Emoji: "💥"},
	{Name: "Aurora", Emoji: "🌌"},
	{Name: "Comet", Emoji: "☄️"},
	{Name: "Sun with Face", Emoji: "🌞"},
	{Name: "Sun Behind Small Cloud", Emoji: "🌤️"},
	{Name: "Sun Behind Large Cloud", Emoji: "🌥️"},
	{Name: "Sun Behind Rain Cloud", Emoji: "🌦️"},
	{Name: "Sun Behind Cloud", Emoji: "⛅"},
	{Name: "Cloud with Rain", Emoji: "🌧️"},
	{Name: "Cloud with Lightning", Emoji: "🌩️"},
	{Name: "Cloud with Tornado", Emoji: "🌪️"},
	{Name: "Fog", Emoji: "🌫️"},
	{Name: "Wind Face", Emoji: "🌬️"},
	{Name: "Cyclone", Emoji: "🌀"},
	{Name: "Rainbow", Emoji: "🌈"},
	{Name: "Umbrella", Emoji: "☂️"},
	{Name: "Umbrella with Rain Drops", Emoji: "☔"},
	{Name: "High Voltage", Emoji: "⚡"},
	{Name: "Snowflake", Emoji: "❄️"},
	{Name: "Snowman", Emoji: "☃️"},
	{Name: "Snowman Without Snow", Emoji: "⛄"},
	{Name: "Cloud", Emoji: "☁️"},
	{Name: "Tornado", Emoji: "🌪️"},
}

var templates = template.Must(template.ParseGlob("index.html"))

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", emojis); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	var results []Emoji
	for _, emoji := range emojis {
		if strings.Contains(strings.ToLower(emoji.Name), strings.ToLower(query)) {
			results = append(results, emoji)
		}
	}
	var emojiResults []string
	for _, result := range results {
		emojiResults = append(emojiResults, result.Emoji)
	}
	if err := json.NewEncoder(w).Encode(emojiResults); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
