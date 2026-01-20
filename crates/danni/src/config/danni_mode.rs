use std::str::FromStr;

use serde::{Deserialize, Serialize};

#[derive(Copy, Clone, Debug, Eq, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum DanniMode {
    Auto,
    Approve,
    SmartApprove,
    Chat,
}

impl FromStr for DanniMode {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "auto" => Ok(DanniMode::Auto),
            "approve" => Ok(DanniMode::Approve),
            "smart_approve" => Ok(DanniMode::SmartApprove),
            "chat" => Ok(DanniMode::Chat),
            _ => Err(format!("invalid mode: {}", s)),
        }
    }
}
