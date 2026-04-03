"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import { API_BASE } from "@/lib/api";

interface HelpButtonProps {
	classroomId: string;
	participantId: string;
	materialStepId: string;
	onRequest?: () => void;
}

export function HelpButton({ classroomId, participantId, materialStepId, onRequest }: HelpButtonProps) {
	const [sending, setSending] = useState(false);
	const [sent, setSent] = useState(false);
	const [error, setError] = useState(false);
	const timerRef = useRef<ReturnType<typeof setTimeout> | null>(null);

	useEffect(() => {
		return () => {
			if (timerRef.current !== null) {
				clearTimeout(timerRef.current);
			}
		};
	}, []);

	const handleClick = useCallback(async () => {
		if (sending || sent) return;

		setSending(true);
		setError(false);
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
				timerRef.current = setTimeout(() => setSent(false), 5000);
			} else {
				setError(true);
				timerRef.current = setTimeout(() => setError(false), 3000);
			}
		} catch {
			setError(true);
			timerRef.current = setTimeout(() => setError(false), 3000);
		} finally {
			setSending(false);
		}
	}, [classroomId, participantId, materialStepId, onRequest, sending, sent]);

	const label = error ? "もういちど" : sent ? "✓ よびました" : "🆘 こまった！";

	return (
		<button
			type="button"
			onClick={() => void handleClick()}
			disabled={sending}
			aria-label={error ? "もう一度助けを呼ぶ" : "助けを呼ぶ"}
			style={{
				width: "100%",
				maxWidth: "400px",
				padding: "24px",
				fontSize: "2rem",
				fontWeight: "bold",
				minHeight: "80px",
				borderRadius: "16px",
				border: `3px solid ${error ? "#ff9800" : "#d32f2f"}`,
				backgroundColor: sent ? "#e8f5e9" : error ? "#fff3e0" : sending ? "#ffebee" : "#d32f2f",
				color: sent ? "#2e7d32" : error ? "#e65100" : "#ffffff",
				cursor: sending ? "wait" : "pointer",
				transition: "all 0.3s",
				display: "block",
				margin: "1rem auto",
			}}
		>
			{label}
		</button>
	);
}
