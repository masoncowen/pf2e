package combat

type ICombatant interface {
    GetName() string
    GetInitiative() int
    SetInitiative(int)
    IsDead() bool
    IsDying() bool
    GetDying() int
    GetHealth() int
    GetMaxHealth() int
    ReduceHealth(int)
    GetAC() int
}
