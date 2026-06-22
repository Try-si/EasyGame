package EasyGame

type Config struct {
	ScreenWidth  int    // largeur de la fenêtre
	ScreenHeight int    // hauteur de la fenêtre
	Title        string // titre de la fenêtre
	Map          string // nom de la map à charger au démarrage

	SpritePath string // chemin vers les sprites
	SoundPath  string // chemin vers les sons
	MapsPath   string // chemin vers les maps
	IAsPath    string // chemin vers les IA
}
