package service

// type PlayerToPlayer struct {
// }

// type PlayerDamage struct {
// 	TotalDamage        float32
// 	HullDamage         float32
// 	ShieldDamage       float32
// 	ExplosionDamage    float32
// 	CritDamage         float32
// 	ThermalDamage      float32
// 	KineticDamage      float32
// 	FriendlyFire       float32
// 	CollisionDamage    float32
// 	IgnoreShieldDamage float32
// }

type Aggregator[A, V any] func(agg A, val V) A

type Filter[T, V any] func(val T) (res V, ok bool)

func Sum(summ float32, val float32) float32 {
	return summ + val
}

func Append[T any, E ~[]T](arr E, val T) E {
	return append(arr, val)
}

// func Average

func ProcessArray[T, A, E any](
	data []T,
	filter Filter[T, E],
	aggregator Aggregator[A, E],
) (res A) {
	for _, line := range data {
		val, ok := filter(line)
		if ok {
			res = aggregator(res, val)
		}
	}
	return res
}
