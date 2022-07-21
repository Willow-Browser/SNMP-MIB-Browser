from unittest.mock import patch, MagicMock
import pytest

from MIB_Browser.scripts.parse_mibs import parse_mibs


@pytest.fixture(autouse=True)
def mock_file_writer():
    with patch("MIB_Browser.scripts.parse_mibs.TestFileWrite") as mock:
        instance = mock.return_value
        instance.setOptions.return_value = instance
        yield mock


@pytest.mark.parametrize("file_name", [("MIB-NAME.txt"), "MIB-NAME"])
def test_parse_mib(file_name):
    with patch(
        "MIB_Browser.scripts.parse_mibs.JsonCodeGen"
    ) as mock_json_code_gen, patch(
        "MIB_Browser.scripts.parse_mibs.StubSearcher"
    ) as mock_stub_searcher, patch(
        "MIB_Browser.scripts.parse_mibs.MibCompiler"
    ) as mock_compiler, patch(
        "MIB_Browser.scripts.parse_mibs.SmiV2Parser"
    ) as mock_parser:
        mocked_base_mibs = MagicMock()

        mock_stub_searcher_instance = mock_stub_searcher.return_value

        mock_json_code_gen_instance = mock_json_code_gen.return_value
        mock_json_code_gen.baseMibs = mocked_base_mibs

        parse_mibs(file_name, MagicMock())
        mock_stub_searcher.assert_called_once_with(*mocked_base_mibs)
    pass
