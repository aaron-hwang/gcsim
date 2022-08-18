package curves

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/player/character/profile"
	"github.com/genshinsim/gcsim/pkg/core/player/weapon"
)

type CharBase struct {
	Rarity     int                `json:"rarity"`
	Body       profile.BodyType   `json:"-"`
	Element    attributes.Element `json:"element"`
	Region     profile.ZoneType   `json:"-"`
	WeaponType weapon.WeaponClass `json:"weapon_class"`

	HPCurve        CharStatCurve   `json:"-"`
	AtkCurve       CharStatCurve   `json:"-"`
	DefCurve       CharStatCurve   `json:"-"`
	BaseHP         float64         `json:"-"`
	BaseAtk        float64         `json:"-"`
	BaseDef        float64         `json:"-"`
	Specialized    attributes.Stat `json:"-"`
	PromotionBonus []PromoData     `json:"-"`
}

type PromoData struct {
	MaxLevel int
	HP       float64
	Atk      float64
	Def      float64
	Special  float64
}

// CharStatGrowthMult provide multiplier for each lvl with 0 being lvl 1 (up to 100)
var CharStatGrowthMult = []map[CharStatCurve]float64{
	{
		GROW_CURVE_HP_S4:     1,
		GROW_CURVE_ATTACK_S4: 1,
		GROW_CURVE_HP_S5:     1,
		GROW_CURVE_ATTACK_S5: 1,
	},
	{
		GROW_CURVE_HP_S4:     1.0829999446868896,
		GROW_CURVE_ATTACK_S4: 1.0829999446868896,
		GROW_CURVE_HP_S5:     1.0829999446868896,
		GROW_CURVE_ATTACK_S5: 1.0829999446868896,
	},
	{
		GROW_CURVE_HP_S4:     1.1649999618530273,
		GROW_CURVE_ATTACK_S4: 1.1649999618530273,
		GROW_CURVE_HP_S5:     1.1660000085830688,
		GROW_CURVE_ATTACK_S5: 1.1660000085830688,
	},
	{
		GROW_CURVE_HP_S4:     1.2480000257492065,
		GROW_CURVE_ATTACK_S4: 1.2480000257492065,
		GROW_CURVE_HP_S5:     1.25,
		GROW_CURVE_ATTACK_S5: 1.25,
	},
	{
		GROW_CURVE_HP_S4:     1.3300000429153442,
		GROW_CURVE_ATTACK_S4: 1.3300000429153442,
		GROW_CURVE_HP_S5:     1.3329999446868896,
		GROW_CURVE_ATTACK_S5: 1.3329999446868896,
	},
	{
		GROW_CURVE_HP_S4:     1.4129999876022339,
		GROW_CURVE_ATTACK_S4: 1.4129999876022339,
		GROW_CURVE_HP_S5:     1.4170000553131104,
		GROW_CURVE_ATTACK_S5: 1.4170000553131104,
	},
	{
		GROW_CURVE_HP_S4:     1.4950000047683716,
		GROW_CURVE_ATTACK_S4: 1.4950000047683716,
		GROW_CURVE_HP_S5:     1.5,
		GROW_CURVE_ATTACK_S5: 1.5,
	},
	{
		GROW_CURVE_HP_S4:     1.5779999494552612,
		GROW_CURVE_ATTACK_S4: 1.5779999494552612,
		GROW_CURVE_HP_S5:     1.5839999914169312,
		GROW_CURVE_ATTACK_S5: 1.5839999914169312,
	},
	{
		GROW_CURVE_HP_S4:     1.6610000133514404,
		GROW_CURVE_ATTACK_S4: 1.6610000133514404,
		GROW_CURVE_HP_S5:     1.6679999828338623,
		GROW_CURVE_ATTACK_S5: 1.6679999828338623,
	},
	{
		GROW_CURVE_HP_S4:     1.7430000305175781,
		GROW_CURVE_ATTACK_S4: 1.7430000305175781,
		GROW_CURVE_HP_S5:     1.7510000467300415,
		GROW_CURVE_ATTACK_S5: 1.7510000467300415,
	},
	{
		GROW_CURVE_HP_S4:     1.8259999752044678,
		GROW_CURVE_ATTACK_S4: 1.8259999752044678,
		GROW_CURVE_HP_S5:     1.8350000381469727,
		GROW_CURVE_ATTACK_S5: 1.8350000381469727,
	},
	{
		GROW_CURVE_HP_S4:     1.9079999923706055,
		GROW_CURVE_ATTACK_S4: 1.9079999923706055,
		GROW_CURVE_HP_S5:     1.9190000295639038,
		GROW_CURVE_ATTACK_S5: 1.9190000295639038,
	},
	{
		GROW_CURVE_HP_S4:     1.9910000562667847,
		GROW_CURVE_ATTACK_S4: 1.9910000562667847,
		GROW_CURVE_HP_S5:     2.003000020980835,
		GROW_CURVE_ATTACK_S5: 2.003000020980835,
	},
	{
		GROW_CURVE_HP_S4:     2.072999954223633,
		GROW_CURVE_ATTACK_S4: 2.072999954223633,
		GROW_CURVE_HP_S5:     2.0880000591278076,
		GROW_CURVE_ATTACK_S5: 2.0880000591278076,
	},
	{
		GROW_CURVE_HP_S4:     2.1559998989105225,
		GROW_CURVE_ATTACK_S4: 2.1559998989105225,
		GROW_CURVE_HP_S5:     2.171999931335449,
		GROW_CURVE_ATTACK_S5: 2.171999931335449,
	},
	{
		GROW_CURVE_HP_S4:     2.239000082015991,
		GROW_CURVE_ATTACK_S4: 2.239000082015991,
		GROW_CURVE_HP_S5:     2.25600004196167,
		GROW_CURVE_ATTACK_S5: 2.25600004196167,
	},
	{
		GROW_CURVE_HP_S4:     2.321000099182129,
		GROW_CURVE_ATTACK_S4: 2.321000099182129,
		GROW_CURVE_HP_S5:     2.3410000801086426,
		GROW_CURVE_ATTACK_S5: 2.3410000801086426,
	},
	{
		GROW_CURVE_HP_S4:     2.4040000438690186,
		GROW_CURVE_ATTACK_S4: 2.4040000438690186,
		GROW_CURVE_HP_S5:     2.424999952316284,
		GROW_CURVE_ATTACK_S5: 2.424999952316284,
	},
	{
		GROW_CURVE_HP_S4:     2.4860000610351562,
		GROW_CURVE_ATTACK_S4: 2.4860000610351562,
		GROW_CURVE_HP_S5:     2.509999990463257,
		GROW_CURVE_ATTACK_S5: 2.509999990463257,
	},
	{
		GROW_CURVE_HP_S4:     2.569000005722046,
		GROW_CURVE_ATTACK_S4: 2.569000005722046,
		GROW_CURVE_HP_S5:     2.5940001010894775,
		GROW_CURVE_ATTACK_S5: 2.5940001010894775,
	},
	{
		GROW_CURVE_HP_S4:     2.6510000228881836,
		GROW_CURVE_ATTACK_S4: 2.6510000228881836,
		GROW_CURVE_HP_S5:     2.678999900817871,
		GROW_CURVE_ATTACK_S5: 2.678999900817871,
	},
	{
		GROW_CURVE_HP_S4:     2.7339999675750732,
		GROW_CURVE_ATTACK_S4: 2.7339999675750732,
		GROW_CURVE_HP_S5:     2.7639999389648438,
		GROW_CURVE_ATTACK_S5: 2.7639999389648438,
	},
	{
		GROW_CURVE_HP_S4:     2.816999912261963,
		GROW_CURVE_ATTACK_S4: 2.816999912261963,
		GROW_CURVE_HP_S5:     2.8489999771118164,
		GROW_CURVE_ATTACK_S5: 2.8489999771118164,
	},
	{
		GROW_CURVE_HP_S4:     2.8989999294281006,
		GROW_CURVE_ATTACK_S4: 2.8989999294281006,
		GROW_CURVE_HP_S5:     2.934000015258789,
		GROW_CURVE_ATTACK_S5: 2.934000015258789,
	},
	{
		GROW_CURVE_HP_S4:     2.9820001125335693,
		GROW_CURVE_ATTACK_S4: 2.9820001125335693,
		GROW_CURVE_HP_S5:     3.0190000534057617,
		GROW_CURVE_ATTACK_S5: 3.0190000534057617,
	},
	{
		GROW_CURVE_HP_S4:     3.063999891281128,
		GROW_CURVE_ATTACK_S4: 3.063999891281128,
		GROW_CURVE_HP_S5:     3.1050000190734863,
		GROW_CURVE_ATTACK_S5: 3.1050000190734863,
	},
	{
		GROW_CURVE_HP_S4:     3.1470000743865967,
		GROW_CURVE_ATTACK_S4: 3.1470000743865967,
		GROW_CURVE_HP_S5:     3.190000057220459,
		GROW_CURVE_ATTACK_S5: 3.190000057220459,
	},
	{
		GROW_CURVE_HP_S4:     3.2290000915527344,
		GROW_CURVE_ATTACK_S4: 3.2290000915527344,
		GROW_CURVE_HP_S5:     3.2750000953674316,
		GROW_CURVE_ATTACK_S5: 3.2750000953674316,
	},
	{
		GROW_CURVE_HP_S4:     3.312000036239624,
		GROW_CURVE_ATTACK_S4: 3.312000036239624,
		GROW_CURVE_HP_S5:     3.3610000610351562,
		GROW_CURVE_ATTACK_S5: 3.3610000610351562,
	},
	{
		GROW_CURVE_HP_S4:     3.3940000534057617,
		GROW_CURVE_ATTACK_S4: 3.3940000534057617,
		GROW_CURVE_HP_S5:     3.446000099182129,
		GROW_CURVE_ATTACK_S5: 3.446000099182129,
	},
	{
		GROW_CURVE_HP_S4:     3.4769999980926514,
		GROW_CURVE_ATTACK_S4: 3.4769999980926514,
		GROW_CURVE_HP_S5:     3.5320000648498535,
		GROW_CURVE_ATTACK_S5: 3.5320000648498535,
	},
	{
		GROW_CURVE_HP_S4:     3.559999942779541,
		GROW_CURVE_ATTACK_S4: 3.559999942779541,
		GROW_CURVE_HP_S5:     3.618000030517578,
		GROW_CURVE_ATTACK_S5: 3.618000030517578,
	},
	{
		GROW_CURVE_HP_S4:     3.6419999599456787,
		GROW_CURVE_ATTACK_S4: 3.6419999599456787,
		GROW_CURVE_HP_S5:     3.7039999961853027,
		GROW_CURVE_ATTACK_S5: 3.7039999961853027,
	},
	{
		GROW_CURVE_HP_S4:     3.7249999046325684,
		GROW_CURVE_ATTACK_S4: 3.7249999046325684,
		GROW_CURVE_HP_S5:     3.7890000343322754,
		GROW_CURVE_ATTACK_S5: 3.7890000343322754,
	},
	{
		GROW_CURVE_HP_S4:     3.806999921798706,
		GROW_CURVE_ATTACK_S4: 3.806999921798706,
		GROW_CURVE_HP_S5:     3.875,
		GROW_CURVE_ATTACK_S5: 3.875,
	},
	{
		GROW_CURVE_HP_S4:     3.890000104904175,
		GROW_CURVE_ATTACK_S4: 3.890000104904175,
		GROW_CURVE_HP_S5:     3.9619998931884766,
		GROW_CURVE_ATTACK_S5: 3.9619998931884766,
	},
	{
		GROW_CURVE_HP_S4:     3.9719998836517334,
		GROW_CURVE_ATTACK_S4: 3.9719998836517334,
		GROW_CURVE_HP_S5:     4.047999858856201,
		GROW_CURVE_ATTACK_S5: 4.047999858856201,
	},
	{
		GROW_CURVE_HP_S4:     4.054999828338623,
		GROW_CURVE_ATTACK_S4: 4.054999828338623,
		GROW_CURVE_HP_S5:     4.133999824523926,
		GROW_CURVE_ATTACK_S5: 4.133999824523926,
	},
	{
		GROW_CURVE_HP_S4:     4.138000011444092,
		GROW_CURVE_ATTACK_S4: 4.138000011444092,
		GROW_CURVE_HP_S5:     4.21999979019165,
		GROW_CURVE_ATTACK_S5: 4.21999979019165,
	},
	{
		GROW_CURVE_HP_S4:     4.21999979019165,
		GROW_CURVE_ATTACK_S4: 4.21999979019165,
		GROW_CURVE_HP_S5:     4.307000160217285,
		GROW_CURVE_ATTACK_S5: 4.307000160217285,
	},
	{
		GROW_CURVE_HP_S4:     4.302999973297119,
		GROW_CURVE_ATTACK_S4: 4.302999973297119,
		GROW_CURVE_HP_S5:     4.39300012588501,
		GROW_CURVE_ATTACK_S5: 4.39300012588501,
	},
	{
		GROW_CURVE_HP_S4:     4.385000228881836,
		GROW_CURVE_ATTACK_S4: 4.385000228881836,
		GROW_CURVE_HP_S5:     4.480000019073486,
		GROW_CURVE_ATTACK_S5: 4.480000019073486,
	},
	{
		GROW_CURVE_HP_S4:     4.4679999351501465,
		GROW_CURVE_ATTACK_S4: 4.4679999351501465,
		GROW_CURVE_HP_S5:     4.566999912261963,
		GROW_CURVE_ATTACK_S5: 4.566999912261963,
	},
	{
		GROW_CURVE_HP_S4:     4.550000190734863,
		GROW_CURVE_ATTACK_S4: 4.550000190734863,
		GROW_CURVE_HP_S5:     4.6529998779296875,
		GROW_CURVE_ATTACK_S5: 4.6529998779296875,
	},
	{
		GROW_CURVE_HP_S4:     4.632999897003174,
		GROW_CURVE_ATTACK_S4: 4.632999897003174,
		GROW_CURVE_HP_S5:     4.739999771118164,
		GROW_CURVE_ATTACK_S5: 4.739999771118164,
	},
	{
		GROW_CURVE_HP_S4:     4.716000080108643,
		GROW_CURVE_ATTACK_S4: 4.716000080108643,
		GROW_CURVE_HP_S5:     4.827000141143799,
		GROW_CURVE_ATTACK_S5: 4.827000141143799,
	},
	{
		GROW_CURVE_HP_S4:     4.797999858856201,
		GROW_CURVE_ATTACK_S4: 4.797999858856201,
		GROW_CURVE_HP_S5:     4.914000034332275,
		GROW_CURVE_ATTACK_S5: 4.914000034332275,
	},
	{
		GROW_CURVE_HP_S4:     4.88100004196167,
		GROW_CURVE_ATTACK_S4: 4.88100004196167,
		GROW_CURVE_HP_S5:     5.000999927520752,
		GROW_CURVE_ATTACK_S5: 5.000999927520752,
	},
	{
		GROW_CURVE_HP_S4:     4.9629998207092285,
		GROW_CURVE_ATTACK_S4: 4.9629998207092285,
		GROW_CURVE_HP_S5:     5.089000225067139,
		GROW_CURVE_ATTACK_S5: 5.089000225067139,
	},
	{
		GROW_CURVE_HP_S4:     5.046000003814697,
		GROW_CURVE_ATTACK_S4: 5.046000003814697,
		GROW_CURVE_HP_S5:     5.176000118255615,
		GROW_CURVE_ATTACK_S5: 5.176000118255615,
	},
	{
		GROW_CURVE_HP_S4:     5.127999782562256,
		GROW_CURVE_ATTACK_S4: 5.127999782562256,
		GROW_CURVE_HP_S5:     5.263000011444092,
		GROW_CURVE_ATTACK_S5: 5.263000011444092,
	},
	{
		GROW_CURVE_HP_S4:     5.210999965667725,
		GROW_CURVE_ATTACK_S4: 5.210999965667725,
		GROW_CURVE_HP_S5:     5.35099983215332,
		GROW_CURVE_ATTACK_S5: 5.35099983215332,
	},
	{
		GROW_CURVE_HP_S4:     5.294000148773193,
		GROW_CURVE_ATTACK_S4: 5.294000148773193,
		GROW_CURVE_HP_S5:     5.438000202178955,
		GROW_CURVE_ATTACK_S5: 5.438000202178955,
	},
	{
		GROW_CURVE_HP_S4:     5.375999927520752,
		GROW_CURVE_ATTACK_S4: 5.375999927520752,
		GROW_CURVE_HP_S5:     5.526000022888184,
		GROW_CURVE_ATTACK_S5: 5.526000022888184,
	},
	{
		GROW_CURVE_HP_S4:     5.459000110626221,
		GROW_CURVE_ATTACK_S4: 5.459000110626221,
		GROW_CURVE_HP_S5:     5.613999843597412,
		GROW_CURVE_ATTACK_S5: 5.613999843597412,
	},
	{
		GROW_CURVE_HP_S4:     5.540999889373779,
		GROW_CURVE_ATTACK_S4: 5.540999889373779,
		GROW_CURVE_HP_S5:     5.702000141143799,
		GROW_CURVE_ATTACK_S5: 5.702000141143799,
	},
	{
		GROW_CURVE_HP_S4:     5.624000072479248,
		GROW_CURVE_ATTACK_S4: 5.624000072479248,
		GROW_CURVE_HP_S5:     5.789999961853027,
		GROW_CURVE_ATTACK_S5: 5.789999961853027,
	},
	{
		GROW_CURVE_HP_S4:     5.705999851226807,
		GROW_CURVE_ATTACK_S4: 5.705999851226807,
		GROW_CURVE_HP_S5:     5.877999782562256,
		GROW_CURVE_ATTACK_S5: 5.877999782562256,
	},
	{
		GROW_CURVE_HP_S4:     5.789000034332275,
		GROW_CURVE_ATTACK_S4: 5.789000034332275,
		GROW_CURVE_HP_S5:     5.966000080108643,
		GROW_CURVE_ATTACK_S5: 5.966000080108643,
	},
	{
		GROW_CURVE_HP_S4:     5.872000217437744,
		GROW_CURVE_ATTACK_S4: 5.872000217437744,
		GROW_CURVE_HP_S5:     6.053999900817871,
		GROW_CURVE_ATTACK_S5: 6.053999900817871,
	},
	{
		GROW_CURVE_HP_S4:     5.953999996185303,
		GROW_CURVE_ATTACK_S4: 5.953999996185303,
		GROW_CURVE_HP_S5:     6.142000198364258,
		GROW_CURVE_ATTACK_S5: 6.142000198364258,
	},
	{
		GROW_CURVE_HP_S4:     6.0370001792907715,
		GROW_CURVE_ATTACK_S4: 6.0370001792907715,
		GROW_CURVE_HP_S5:     6.230000019073486,
		GROW_CURVE_ATTACK_S5: 6.230000019073486,
	},
	{
		GROW_CURVE_HP_S4:     6.11899995803833,
		GROW_CURVE_ATTACK_S4: 6.11899995803833,
		GROW_CURVE_HP_S5:     6.318999767303467,
		GROW_CURVE_ATTACK_S5: 6.318999767303467,
	},
	{
		GROW_CURVE_HP_S4:     6.202000141143799,
		GROW_CURVE_ATTACK_S4: 6.202000141143799,
		GROW_CURVE_HP_S5:     6.4070000648498535,
		GROW_CURVE_ATTACK_S5: 6.4070000648498535,
	},
	{
		GROW_CURVE_HP_S4:     6.283999919891357,
		GROW_CURVE_ATTACK_S4: 6.283999919891357,
		GROW_CURVE_HP_S5:     6.495999813079834,
		GROW_CURVE_ATTACK_S5: 6.495999813079834,
	},
	{
		GROW_CURVE_HP_S4:     6.367000102996826,
		GROW_CURVE_ATTACK_S4: 6.367000102996826,
		GROW_CURVE_HP_S5:     6.585000038146973,
		GROW_CURVE_ATTACK_S5: 6.585000038146973,
	},
	{
		GROW_CURVE_HP_S4:     6.449999809265137,
		GROW_CURVE_ATTACK_S4: 6.449999809265137,
		GROW_CURVE_HP_S5:     6.672999858856201,
		GROW_CURVE_ATTACK_S5: 6.672999858856201,
	},
	{
		GROW_CURVE_HP_S4:     6.5320000648498535,
		GROW_CURVE_ATTACK_S4: 6.5320000648498535,
		GROW_CURVE_HP_S5:     6.76200008392334,
		GROW_CURVE_ATTACK_S5: 6.76200008392334,
	},
	{
		GROW_CURVE_HP_S4:     6.614999771118164,
		GROW_CURVE_ATTACK_S4: 6.614999771118164,
		GROW_CURVE_HP_S5:     6.85099983215332,
		GROW_CURVE_ATTACK_S5: 6.85099983215332,
	},
	{
		GROW_CURVE_HP_S4:     6.697000026702881,
		GROW_CURVE_ATTACK_S4: 6.697000026702881,
		GROW_CURVE_HP_S5:     6.940000057220459,
		GROW_CURVE_ATTACK_S5: 6.940000057220459,
	},
	{
		GROW_CURVE_HP_S4:     6.78000020980835,
		GROW_CURVE_ATTACK_S4: 6.78000020980835,
		GROW_CURVE_HP_S5:     7.0289998054504395,
		GROW_CURVE_ATTACK_S5: 7.0289998054504395,
	},
	{
		GROW_CURVE_HP_S4:     6.861999988555908,
		GROW_CURVE_ATTACK_S4: 6.861999988555908,
		GROW_CURVE_HP_S5:     7.11899995803833,
		GROW_CURVE_ATTACK_S5: 7.11899995803833,
	},
	{
		GROW_CURVE_HP_S4:     6.945000171661377,
		GROW_CURVE_ATTACK_S4: 6.945000171661377,
		GROW_CURVE_HP_S5:     7.208000183105469,
		GROW_CURVE_ATTACK_S5: 7.208000183105469,
	},
	{
		GROW_CURVE_HP_S4:     7.0279998779296875,
		GROW_CURVE_ATTACK_S4: 7.0279998779296875,
		GROW_CURVE_HP_S5:     7.296999931335449,
		GROW_CURVE_ATTACK_S5: 7.296999931335449,
	},
	{
		GROW_CURVE_HP_S4:     7.110000133514404,
		GROW_CURVE_ATTACK_S4: 7.110000133514404,
		GROW_CURVE_HP_S5:     7.38700008392334,
		GROW_CURVE_ATTACK_S5: 7.38700008392334,
	},
	{
		GROW_CURVE_HP_S4:     7.192999839782715,
		GROW_CURVE_ATTACK_S4: 7.192999839782715,
		GROW_CURVE_HP_S5:     7.47599983215332,
		GROW_CURVE_ATTACK_S5: 7.47599983215332,
	},
	{
		GROW_CURVE_HP_S4:     7.275000095367432,
		GROW_CURVE_ATTACK_S4: 7.275000095367432,
		GROW_CURVE_HP_S5:     7.565999984741211,
		GROW_CURVE_ATTACK_S5: 7.565999984741211,
	},
	{
		GROW_CURVE_HP_S4:     7.357999801635742,
		GROW_CURVE_ATTACK_S4: 7.357999801635742,
		GROW_CURVE_HP_S5:     7.656000137329102,
		GROW_CURVE_ATTACK_S5: 7.656000137329102,
	},
	{
		GROW_CURVE_HP_S4:     7.440000057220459,
		GROW_CURVE_ATTACK_S4: 7.440000057220459,
		GROW_CURVE_HP_S5:     7.745999813079834,
		GROW_CURVE_ATTACK_S5: 7.745999813079834,
	},
	{
		GROW_CURVE_HP_S4:     7.5229997634887695,
		GROW_CURVE_ATTACK_S4: 7.5229997634887695,
		GROW_CURVE_HP_S5:     7.835999965667725,
		GROW_CURVE_ATTACK_S5: 7.835999965667725,
	},
	{
		GROW_CURVE_HP_S4:     7.605999946594238,
		GROW_CURVE_ATTACK_S4: 7.605999946594238,
		GROW_CURVE_HP_S5:     7.926000118255615,
		GROW_CURVE_ATTACK_S5: 7.926000118255615,
	},
	{
		GROW_CURVE_HP_S4:     7.688000202178955,
		GROW_CURVE_ATTACK_S4: 7.688000202178955,
		GROW_CURVE_HP_S5:     8.015999794006348,
		GROW_CURVE_ATTACK_S5: 8.015999794006348,
	},
	{
		GROW_CURVE_HP_S4:     7.770999908447266,
		GROW_CURVE_ATTACK_S4: 7.770999908447266,
		GROW_CURVE_HP_S5:     8.105999946594238,
		GROW_CURVE_ATTACK_S5: 8.105999946594238,
	},
	{
		GROW_CURVE_HP_S4:     7.853000164031982,
		GROW_CURVE_ATTACK_S4: 7.853000164031982,
		GROW_CURVE_HP_S5:     8.196000099182129,
		GROW_CURVE_ATTACK_S5: 8.196000099182129,
	},
	{
		GROW_CURVE_HP_S4:     7.935999870300293,
		GROW_CURVE_ATTACK_S4: 7.935999870300293,
		GROW_CURVE_HP_S5:     8.28600025177002,
		GROW_CURVE_ATTACK_S5: 8.28600025177002,
	},
	{
		GROW_CURVE_HP_S4:     8.017999649047852,
		GROW_CURVE_ATTACK_S4: 8.017999649047852,
		GROW_CURVE_HP_S5:     8.376999855041504,
		GROW_CURVE_ATTACK_S5: 8.376999855041504,
	},
	{
		GROW_CURVE_HP_S4:     8.10099983215332,
		GROW_CURVE_ATTACK_S4: 8.10099983215332,
		GROW_CURVE_HP_S5:     8.467000007629395,
		GROW_CURVE_ATTACK_S5: 8.467000007629395,
	},
	{
		GROW_CURVE_HP_S4:     8.182999610900879,
		GROW_CURVE_ATTACK_S4: 8.182999610900879,
		GROW_CURVE_HP_S5:     8.557999610900879,
		GROW_CURVE_ATTACK_S5: 8.557999610900879,
	},
	{
		GROW_CURVE_HP_S4:     8.265999794006348,
		GROW_CURVE_ATTACK_S4: 8.265999794006348,
		GROW_CURVE_HP_S5:     8.64900016784668,
		GROW_CURVE_ATTACK_S5: 8.64900016784668,
	},
	{
		GROW_CURVE_HP_S4:     8.348999977111816,
		GROW_CURVE_ATTACK_S4: 8.348999977111816,
		GROW_CURVE_HP_S5:     8.73900032043457,
		GROW_CURVE_ATTACK_S5: 8.73900032043457,
	},
	{
		GROW_CURVE_HP_S4:     8.430999755859375,
		GROW_CURVE_ATTACK_S4: 8.430999755859375,
		GROW_CURVE_HP_S5:     8.829999923706055,
		GROW_CURVE_ATTACK_S5: 8.829999923706055,
	},
	{
		GROW_CURVE_HP_S4:     8.513999938964844,
		GROW_CURVE_ATTACK_S4: 8.513999938964844,
		GROW_CURVE_HP_S5:     8.920999526977539,
		GROW_CURVE_ATTACK_S5: 8.920999526977539,
	},
	{
		GROW_CURVE_HP_S4:     8.595999717712402,
		GROW_CURVE_ATTACK_S4: 8.595999717712402,
		GROW_CURVE_HP_S5:     9.01200008392334,
		GROW_CURVE_ATTACK_S5: 9.01200008392334,
	},
	{
		GROW_CURVE_HP_S4:     8.678999900817871,
		GROW_CURVE_ATTACK_S4: 8.678999900817871,
		GROW_CURVE_HP_S5:     9.102999687194824,
		GROW_CURVE_ATTACK_S5: 9.102999687194824,
	},
	{
		GROW_CURVE_HP_S4:     8.76099967956543,
		GROW_CURVE_ATTACK_S4: 8.76099967956543,
		GROW_CURVE_HP_S5:     9.194999694824219,
		GROW_CURVE_ATTACK_S5: 9.194999694824219,
	},
	{
		GROW_CURVE_HP_S4:     8.843999862670898,
		GROW_CURVE_ATTACK_S4: 8.843999862670898,
		GROW_CURVE_HP_S5:     9.28600025177002,
		GROW_CURVE_ATTACK_S5: 9.28600025177002,
	},
	{
		GROW_CURVE_HP_S4:     8.927000045776367,
		GROW_CURVE_ATTACK_S4: 8.927000045776367,
		GROW_CURVE_HP_S5:     9.376999855041504,
		GROW_CURVE_ATTACK_S5: 9.376999855041504,
	},
	{
		GROW_CURVE_HP_S4:     9.008999824523926,
		GROW_CURVE_ATTACK_S4: 9.008999824523926,
		GROW_CURVE_HP_S5:     9.468999862670898,
		GROW_CURVE_ATTACK_S5: 9.468999862670898,
	},
	{
		GROW_CURVE_HP_S4:     9.092000007629395,
		GROW_CURVE_ATTACK_S4: 9.092000007629395,
		GROW_CURVE_HP_S5:     9.5600004196167,
		GROW_CURVE_ATTACK_S5: 9.5600004196167,
	},
	{
		GROW_CURVE_HP_S4:     9.173999786376953,
		GROW_CURVE_ATTACK_S4: 9.173999786376953,
		GROW_CURVE_HP_S5:     9.652000427246094,
		GROW_CURVE_ATTACK_S5: 9.652000427246094,
	},
}
