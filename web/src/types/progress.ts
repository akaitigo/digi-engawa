export interface LearnerProgress {
	id: string;
	participant_id: string;
	material_id: string;
	current_step: number;
	completed: boolean;
	updated_at: string;
}
