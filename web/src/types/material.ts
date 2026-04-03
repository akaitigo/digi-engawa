export interface Material {
	id: string;
	title: string;
	description: string;
	steps: Step[];
	created_at: string;
	updated_at: string;
}

export interface Step {
	id: string;
	material_id: string;
	step_order: number;
	title: string;
	body: string;
	furigana_body: string;
	audio_text: string;
	created_at: string;
}
