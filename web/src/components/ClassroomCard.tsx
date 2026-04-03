"use client";

import type { Classroom } from "@/types/classroom";

interface ClassroomCardProps {
	classroom: Classroom;
	onClick?: () => void;
}

export function ClassroomCard({ classroom, onClick }: ClassroomCardProps) {
	const date = new Date(classroom.scheduled_at);
	const dateStr = date.toLocaleDateString("ja-JP", {
		year: "numeric",
		month: "long",
		day: "numeric",
		weekday: "short",
	});
	const timeStr = date.toLocaleTimeString("ja-JP", {
		hour: "2-digit",
		minute: "2-digit",
	});

	return (
		<button
			type="button"
			onClick={onClick}
			style={{
				width: "100%",
				padding: "1.25rem",
				borderRadius: "12px",
				border: "2px solid #e0e0e0",
				backgroundColor: "#ffffff",
				textAlign: "left",
				cursor: onClick ? "pointer" : "default",
				marginBottom: "0.75rem",
				display: "block",
			}}
		>
			<h3 style={{ fontSize: "1.375rem", fontWeight: "bold", marginBottom: "0.5rem", color: "#333" }}>
				{classroom.title}
			</h3>
			<p style={{ fontSize: "1rem", color: "#666", marginBottom: "0.25rem" }}>
				{dateStr} {timeStr}
			</p>
			<p style={{ fontSize: "1rem", color: "#666", marginBottom: "0.25rem" }}>{classroom.location}</p>
			<div style={{ display: "flex", justifyContent: "space-between", marginTop: "0.5rem" }}>
				<span
					style={{
						fontSize: "0.875rem",
						padding: "4px 8px",
						borderRadius: "8px",
						backgroundColor: "#e8f5e9",
						color: "#2e7d32",
					}}
				>
					コード: {classroom.classroom_code}
				</span>
				<span style={{ fontSize: "0.875rem", color: "#999" }}>定員 {String(classroom.capacity)}名</span>
			</div>
		</button>
	);
}
