export interface Classroom {
	id: string;
	title: string;
	description: string;
	location: string;
	capacity: number;
	scheduled_at: string;
	classroom_code: string;
	created_at: string;
	updated_at: string;
}

export interface Participant {
	id: string;
	classroom_id: string;
	name: string;
	role: "learner" | "supporter" | "organizer";
	created_at: string;
}
