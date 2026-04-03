"use client";

import type { Participant } from "@/types/classroom";

interface ParticipantListProps {
	participants: Participant[];
}

const roleLabels: Record<string, string> = {
	learner: "受講者",
	supporter: "サポーター",
	organizer: "主催者",
};

const roleColors: Record<string, string> = {
	learner: "#e3f2fd",
	supporter: "#e8f5e9",
	organizer: "#fff3e0",
};

export function ParticipantList({ participants }: ParticipantListProps) {
	if (participants.length === 0) {
		return <p style={{ color: "#999", textAlign: "center", padding: "1rem" }}>参加者はまだいません</p>;
	}

	return (
		<div>
			{participants.map((p) => (
				<div
					key={p.id}
					style={{
						display: "flex",
						justifyContent: "space-between",
						alignItems: "center",
						padding: "0.75rem 1rem",
						borderRadius: "8px",
						backgroundColor: roleColors[p.role] ?? "#f5f5f5",
						marginBottom: "0.5rem",
					}}
				>
					<span style={{ fontSize: "1.125rem", fontWeight: "bold" }}>{p.name}</span>
					<span
						style={{
							fontSize: "0.875rem",
							padding: "2px 8px",
							borderRadius: "12px",
							backgroundColor: "#ffffff",
						}}
					>
						{roleLabels[p.role] ?? p.role}
					</span>
				</div>
			))}
		</div>
	);
}
