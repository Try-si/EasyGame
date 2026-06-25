package EasyGame

type config struct {
	ScreenWidth  int    // largeur de la fenêtre
	ScreenHeight int    // hauteur de la fenêtre
	Title        string // titre de la fenêtre
	Map          string // nom de la map à charger au démarrage

	SpritePath []string // chemin vers les sprites
	SoundPath  string   // chemin vers les sons
	MapsPath   string   // chemin vers les maps
	InputsPath string   // chemin vers les inputs
}

type camera struct {
}
type entitys struct {
}
type sound struct {
}
type maps struct {
}
type inputs struct {
}
type sprite struct {
}
type ias struct {
}

var (
	Cam      camera
	Entities entitys
	Sounds   sound
	Maps     maps
	Input    inputs
	Sprites  sprite
	IAss     ias
)

type BehaviourNode struct {
	Condition func(WorldState) (bool, any)
	TrueNode  *BehaviourNode
	FalseNode *BehaviourNode
	Action    func(string, any)
}

type WorldState struct {
}

type Brain struct {
	Transition map[string][]Transition
	States     map[string]State
	Entity     map[string]BrainEntity
}

type Transition struct {
	Weight    float32
	To        string
	Condition func(WorldState) (bool, any)
}

type State struct {
	Action func(string, any)
}

type BrainEntity struct {
	CurrentState string
	PréTrans     Transition
}

type FSM struct {
	Transition map[string][]FSMTransition
	State      map[string]State
	Entity     map[string]FSMEntity
}

type FSMTransition struct {
	To        string
	Condition func(WorldState) (bool, any)
}

type FSMEntity struct {
	CurrentState string
	PréTrans     FSMTransition
}
