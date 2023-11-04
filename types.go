package main

type ProjectAI struct {
	Success             bool   `json:"success"`
	Predictions         []any  `json:"predictions"`
	Message             string `json:"message"`
	ProcessMs           int    `json:"processMs"`
	InferenceMs         int    `json:"inferenceMs"`
	Code                int    `json:"code"`
	Command             string `json:"command"`
	ModuleID            string `json:"moduleId"`
	ExecutionProvider   string `json:"executionProvider"`
	AnalysisRoundTripMs int    `json:"analysisRoundTripMs"`
}

type FrigateEvent struct {
	Before struct {
		ID              string  `json:"id"`
		Camera          string  `json:"camera"`
		FrameTime       float64 `json:"frame_time"`
		SnapshotTime    float64 `json:"snapshot_time"`
		Label           string  `json:"label"`
		SubLabel        any     `json:"sub_label"`
		TopScore        float64 `json:"top_score"`
		FalsePositive   bool    `json:"false_positive"`
		StartTime       float64 `json:"start_time"`
		EndTime         any     `json:"end_time"`
		Score           float64 `json:"score"`
		Box             []int   `json:"box"`
		Area            int     `json:"area"`
		Ratio           float64 `json:"ratio"`
		Region          []int   `json:"region"`
		Stationary      bool    `json:"stationary"`
		MotionlessCount int     `json:"motionless_count"`
		PositionChanges int     `json:"position_changes"`
		CurrentZones    []any   `json:"current_zones"`
		EnteredZones    []any   `json:"entered_zones"`
		HasClip         bool    `json:"has_clip"`
		HasSnapshot     bool    `json:"has_snapshot"`
	} `json:"before"`
	After struct {
		ID              string  `json:"id"`
		Camera          string  `json:"camera"`
		FrameTime       float64 `json:"frame_time"`
		SnapshotTime    float64 `json:"snapshot_time"`
		Label           string  `json:"label"`
		SubLabel        any     `json:"sub_label"`
		TopScore        float64 `json:"top_score"`
		FalsePositive   bool    `json:"false_positive"`
		StartTime       float64 `json:"start_time"`
		EndTime         any     `json:"end_time"`
		Score           float64 `json:"score"`
		Box             []int   `json:"box"`
		Area            int     `json:"area"`
		Ratio           float64 `json:"ratio"`
		Region          []int   `json:"region"`
		Stationary      bool    `json:"stationary"`
		MotionlessCount int     `json:"motionless_count"`
		PositionChanges int     `json:"position_changes"`
		CurrentZones    []any   `json:"current_zones"`
		EnteredZones    []any   `json:"entered_zones"`
		HasClip         bool    `json:"has_clip"`
		HasSnapshot     bool    `json:"has_snapshot"`
	} `json:"after"`
	Type string `json:"type"`
}
