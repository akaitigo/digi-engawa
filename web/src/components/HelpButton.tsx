"use client";

import { useCallback, useState } from "react";

interface HelpButtonProps {
	classroomId: string;
	participantId: string;
	materialStepId: string;
	onRequest?: () => void;
}

const API_BASE = process.env["NEXT_PUBLIC_API_URL"] ?? "http://localhost:8080";

export function HelpButton({ classroomId, participantId, materialStepId, onRequest }: HelpButtonProps) {
	const [sending, setSending] = useState(false);
	const [sent, setSent] = useState(false);

	const handleClick = useCallback(async () => {
		if (sending || sent) return;

		setSending(true);
		try {
			const res = await fetch(`${API_BASE}/api/help-requests`, {
				method: "POST",
				headers: { "Content-Type": "application/json" },
				body: JSON.stringify({
					classroom_id: classroomId,
					participant_id: participantId,
					material_step_id: materialStepId,
				}),
			});

			if (res.ok) {
				setSent(true);
				onRequest?.();
				setTimeout(() => setSent(false), 5000);
			}
		} finally {
			setSending(false);
		}
	}, [classroomId, participantId, materialStepId, sending, sent, onRequest]);

	return (
		<button
			type="button"
			onClick={() => void handleClick()}
			disabled={sending}
			aria-label="助けを呼ぶ"
			style={{
				width: "100%",
				maxWidth: "400px",
				padding: "24px",
				fontSize: "2rem",
				fontWeight: "bold",
				minHeight: "80px",
				borderRadius: "16px",
				border: "3px solid #d32f2f",
				backgroundColor: sent ? "#e8f5e9" : sending ? "#ffebee" : "#d32f2f",
				color: sent ? "#2e7d32" : "#ffffff",
				cursor: sending ? "wait" : "pointer",
				transition: "all 0.3s",
				display: "block",
				margin: "1rem auto",
			}}
		>
			{sent ? "✓ よびました" : "🆘 こまった！"}
		</button>
	);
}
