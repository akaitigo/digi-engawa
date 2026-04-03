"use client";

import { useCallback, useState } from "react";
import type { ClassroomFormData } from "@/components/ClassroomForm";
import { ClassroomForm } from "@/components/ClassroomForm";
import { API_BASE } from "@/lib/api";

export default function NewClassroomPage() {
	const [submitting, setSubmitting] = useState(false);
	const [error, setError] = useState<string | null>(null);

	const handleSubmit = useCallback(async (data: ClassroomFormData) => {
		setSubmitting(true);
		setError(null);
		try {
			const res = await fetch(`${API_BASE}/api/classrooms`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify(data),
			});

			if (res.ok) {
				window.location.href = "/dashboard";
			} else {
				setError("教室の作成に失敗しました");
			}
		} catch {
			setError("ネットワークエラーが発生しました");
		} finally {
			setSubmitting(false);
		}
	}, []);

	return (
		<main style={{ padding: "1rem" }}>
			<h1 style={{ fontSize: "2rem", fontWeight: "bold", textAlign: "center", marginBottom: "1.5rem" }}>
				教室を作成する
			</h1>
			{error && (
				<p
					role="alert"
					style={{
						textAlign: "center",
						color: "#d32f2f",
						fontSize: "1.125rem",
						marginBottom: "1rem",
					}}
				>
					{error}
				</p>
			)}
			<ClassroomForm onSubmit={(data) => void handleSubmit(data)} submitting={submitting} />
		</main>
	);
}
