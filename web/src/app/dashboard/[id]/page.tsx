"use client";

import { use, useEffect, useState } from "react";
import { ParticipantList } from "@/components/ParticipantList";
import type { Classroom, Participant } from "@/types/classroom";

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export default function ClassroomDetailPage({ params }: { params: Promise<{ id: string }> }) {
	const { id } = use(params);
	const [classroom, setClassroom] = useState<Classroom | null>(null);
	const [participants, setParticipants] = useState<Participant[]>([]);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		const load = async () => {
			try {
				const [classroomRes, participantsRes] = await Promise.all([
					fetch(`${API_BASE}/api/classrooms/${id}`),
					fetch(`${API_BASE}/api/classrooms/${id}/participants`),
				]);

				if (classroomRes.ok) {
					setClassroom((await classroomRes.json()) as Classroom);
				}
				if (participantsRes.ok) {
					setParticipants((await participantsRes.json()) as Participant[]);
				}
			} finally {
				setLoading(false);
			}
		};
		void load();
	}, [id]);

	if (loading) {
		return (
			<main style={{ padding: "2rem", textAlign: "center" }}>
				<p style={{ fontSize: "1.5rem" }}>よみこみちゅう...</p>
			</main>
		);
	}

	if (!classroom) {
		return (
			<main style={{ padding: "2rem", textAlign: "center" }}>
				<p style={{ fontSize: "1.5rem", color: "#d32f2f" }}>教室が見つかりません</p>
			</main>
		);
	}

	return (
		<main style={{ maxWidth: "600px", margin: "0 auto", padding: "1rem" }}>
			<h1 style={{ fontSize: "2rem", fontWeight: "bold", marginBottom: "0.5rem" }}>{classroom.title}</h1>
			<p style={{ fontSize: "1rem", color: "#666", marginBottom: "1rem" }}>{classroom.description}</p>

			<div
				style={{
					padding: "1rem",
					borderRadius: "12px",
					backgroundColor: "#f5f5f5",
					marginBottom: "1.5rem",
				}}
			>
				<p style={{ fontSize: "1rem", marginBottom: "0.5rem" }}>
					<strong>場所:</strong> {classroom.location}
				</p>
				<p style={{ fontSize: "1rem", marginBottom: "0.5rem" }}>
					<strong>日時:</strong> {new Date(classroom.scheduled_at).toLocaleString("ja-JP")}
				</p>
				<p style={{ fontSize: "1rem", marginBottom: "0.5rem" }}>
					<strong>定員:</strong> {String(classroom.capacity)}名
				</p>
				<p style={{ fontSize: "1.25rem", fontWeight: "bold" }}>
					教室コード: <span style={{ color: "#4A7C59", letterSpacing: "4px" }}>{classroom.classroom_code}</span>
				</p>
			</div>

			<h2 style={{ fontSize: "1.5rem", fontWeight: "bold", marginBottom: "1rem" }}>
				参加者 ({String(participants.length)})
			</h2>
			<ParticipantList participants={participants} />
		</main>
	);
}
