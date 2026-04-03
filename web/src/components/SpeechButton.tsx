"use client";

import { useCallback, useState } from "react";

interface SpeechButtonProps {
	text: string;
	label?: string;
}

export function SpeechButton({ text, label = "よみあげ" }: SpeechButtonProps) {
	const [speaking, setSpeaking] = useState(false);

	const handleSpeak = useCallback(() => {
		if (typeof window === "undefined" || !window.speechSynthesis) {
			return;
		}

		if (speaking) {
			window.speechSynthesis.cancel();
			setSpeaking(false);
			return;
		}

		const utterance = new SpeechSynthesisUtterance(text);
		utterance.lang = "ja-JP";
		utterance.rate = 0.8;
		utterance.pitch = 1.0;

		utterance.onend = () => setSpeaking(false);
		utterance.onerror = () => setSpeaking(false);

		setSpeaking(true);
		window.speechSynthesis.speak(utterance);
	}, [text, speaking]);

	return (
		<button
			type="button"
			onClick={handleSpeak}
			aria-label={speaking ? "読み上げを停止" : label}
			style={{
				padding: "16px 32px",
				fontSize: "1.25rem",
				fontWeight: "bold",
				minWidth: "48px",
				minHeight: "48px",
				borderRadius: "12px",
				border: "2px solid #4A7C59",
				backgroundColor: speaking ? "#4A7C59" : "#ffffff",
				color: speaking ? "#ffffff" : "#4A7C59",
				cursor: "pointer",
				transition: "all 0.2s",
			}}
		>
			{speaking ? "⏹ とめる" : `🔊 ${label}`}
		</button>
	);
}
