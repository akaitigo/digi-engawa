"use client";

import { useCallback, useState } from "react";
import type { Step } from "@/types/material";
import { FuriganaText } from "./FuriganaText";
import { ProgressBar } from "./ProgressBar";
import { SpeechButton } from "./SpeechButton";

interface StepViewerProps {
	steps: Step[];
	onStepChange?: (stepOrder: number) => void;
}

export function StepViewer({ steps, onStepChange }: StepViewerProps) {
	const [currentIndex, setCurrentIndex] = useState(0);
	const totalSteps = steps.length;

	const currentStep = steps[currentIndex];

	const goToStep = useCallback(
		(index: number) => {
			setCurrentIndex(index);
			const step = steps[index];
			if (step && onStepChange) {
				onStepChange(step.step_order);
			}
		},
		[steps, onStepChange],
	);

	const handlePrev = useCallback(() => {
		if (currentIndex > 0) {
			goToStep(currentIndex - 1);
		}
	}, [currentIndex, goToStep]);

	const handleNext = useCallback(() => {
		if (currentIndex < totalSteps - 1) {
			goToStep(currentIndex + 1);
		}
	}, [currentIndex, totalSteps, goToStep]);

	if (!currentStep) {
		return (
			<div style={{ padding: "2rem", textAlign: "center", fontSize: "1.25rem" }}>
				<p>教材にステップがありません</p>
			</div>
		);
	}

	return (
		<div style={{ maxWidth: "600px", margin: "0 auto", padding: "1rem" }}>
			<ProgressBar current={currentIndex + 1} total={totalSteps} />

			<div
				style={{
					backgroundColor: "#f9f9f9",
					borderRadius: "16px",
					padding: "2rem",
					marginBottom: "1.5rem",
					minHeight: "200px",
				}}
			>
				<h2
					style={{
						fontSize: "1.75rem",
						fontWeight: "bold",
						marginBottom: "1rem",
						color: "#333",
					}}
				>
					{currentStep.title}
				</h2>

				<div style={{ marginBottom: "1.5rem" }}>
					<FuriganaText text={currentStep.body} furigana={currentStep.furigana_body} />
				</div>

				{currentStep.audio_text && (
					<div style={{ textAlign: "center" }}>
						<SpeechButton text={currentStep.audio_text} />
					</div>
				)}
			</div>

			<div
				style={{
					display: "flex",
					justifyContent: "space-between",
					gap: "1rem",
				}}
			>
				<button
					type="button"
					onClick={handlePrev}
					disabled={currentIndex === 0}
					aria-label="前のステップ"
					style={{
						flex: 1,
						padding: "20px",
						fontSize: "1.5rem",
						fontWeight: "bold",
						minHeight: "64px",
						borderRadius: "12px",
						border: "2px solid #ccc",
						backgroundColor: currentIndex === 0 ? "#f0f0f0" : "#ffffff",
						color: currentIndex === 0 ? "#aaa" : "#333",
						cursor: currentIndex === 0 ? "not-allowed" : "pointer",
					}}
				>
					← まえ
				</button>

				<button
					type="button"
					onClick={handleNext}
					disabled={currentIndex === totalSteps - 1}
					aria-label="次のステップ"
					style={{
						flex: 1,
						padding: "20px",
						fontSize: "1.5rem",
						fontWeight: "bold",
						minHeight: "64px",
						borderRadius: "12px",
						border: "2px solid #4A7C59",
						backgroundColor: currentIndex === totalSteps - 1 ? "#f0f0f0" : "#4A7C59",
						color: currentIndex === totalSteps - 1 ? "#aaa" : "#ffffff",
						cursor: currentIndex === totalSteps - 1 ? "not-allowed" : "pointer",
					}}
				>
					つぎ →
				</button>
			</div>
		</div>
	);
}
