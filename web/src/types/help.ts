export interface HelpRequest {
	id: string;
	classroom_id: string;
	participant_id: string;
	material_step_id: string;
	status: "pending" | "in_progress" | "resolved";
	created_at: string;
	resolved_at?: string;
}

export interface WsMessage {
	type: string;
	data: unknown;
}
