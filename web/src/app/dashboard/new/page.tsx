"use client";

import { useCallback, useState } from "react";
import type { ClassroomFormData } from "@/components/ClassroomForm";
import { ClassroomForm } from "@/components/ClassroomForm";

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export default function NewClassroomPage() {
	const [submitting, setSubmitting] = useState(false);

	const handleSubmit = useCallback(async (data: ClassroomFormData) => {
		setSubmitting(true);
		try {
			const res = await fetch(`${API_BASE}/api/classrooms`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(data),
			});

			if (res.ok) {
				window.location.href = "/dashboard";
			}
		} finally {
			setSubmitting(false);
		}
	}, []);

	return (
		<main style={{ padding: "1rem" }}>
			<h1 style={{ fontSize: "2rem", fontWeight: "bold", textAlign: "center", marginBottom: "1.5rem" }}>
				教室を作成する
			</h1>
			<ClassroomForm onSubmit={(data) => void handleSubmit(data)} submitting={submitting} />
		</main>
	);
}
