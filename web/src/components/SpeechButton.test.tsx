import { render, screen, fireEvent } from "@testing-library/react";
import { describe, expect, it, vi, beforeEach } from "vitest";
import { SpeechButton } from "./SpeechButton";

describe("SpeechButton", () => {
	const mockSpeak = vi.fn();
	const mockCancel = vi.fn();

	beforeEach(() => {
		vi.clearAllMocks();

		// Mock SpeechSynthesisUtterance
		const MockUtterance = vi.fn().mockImplementation(() => ({
			lang: "",
			rate: 1,
			pitch: 1,
			onend: null,
			onerror: null,
		}));
		vi.stubGlobal("SpeechSynthesisUtterance", MockUtterance);

		Object.defineProperty(window, "speechSynthesis", {
			value: {
				speak: mockSpeak,
				cancel: mockCancel,
			},
			writable: true,
			configurable: true,
		});
	});

	it("renders with default label", () => {
		render(<SpeechButton text="テスト" />);
		expect(screen.getByText("🔊 よみあげ")).toBeDefined();
	});

	it("renders with custom label", () => {
		render(<SpeechButton text="テスト" label="きく" />);
		expect(screen.getByText("🔊 きく")).toBeDefined();
	});

	it("calls speechSynthesis.speak on click", () => {
		render(<SpeechButton text="テスト音声" />);
		fireEvent.click(screen.getByRole("button"));
		expect(mockSpeak).toHaveBeenCalledTimes(1);
	});

	it("has minimum touch target size", () => {
		render(<SpeechButton text="テスト" />);
		const button = screen.getByRole("button");
		expect(button.style.minWidth).toBe("48px");
		expect(button.style.minHeight).toBe("48px");
	});
});
